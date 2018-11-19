"""
    Switch message from dataprovider to Worker.
    Note: all messages for same hash_key MUST send to same Worker

    Note: all messages for same instance id MUST send to same Switch
"""

from gevent import monkey

monkey.patch_all()
# monkey.patch_ssl()
# monkey.patch_time()
import gevent
import sys
import zmq
import include,random
import traceback
import logging
import os,json
import time,datetime
from collections import deque
from utils.msgpackhelper import unpack_msg, pack_msg
from confintf.data.btr_key import BTRKEY_TS, BTRKEY_DP_SYNC,BTRKEY_IP_SRC,BTRKEY_IP_DST
from confintf.data.db import DECODE_ID
from utils.zeromq import ZMQInputChannel
from scripts.transform import decode_input,make_hash_id,transforme,lookup_transformer,lookup_protocol
from confintf.mongo import get_instance_id_collection
from confintf.datapath import get_cbms_new_app_datapath
from django.conf import settings
from datastore import redis_operation
from confintf.data.db import DKEY_BIZ_TRANSACTION_ID, REVERSE_INTF
from  multiprocessing import Process,Queue
import threading
ZEROMQ_DEFAULT_HWM = getattr(settings, 'ZEROMQ_DEFAULT_HWM', 100000)
logger = logging.getLogger('default')
tc_logger = logging.getLogger('cbms.tc')
DEBUG=False
if  not DEBUG:
    logger.info=str

def cmp_time(now, time_ranges):
    now_sec = now[0] * 60 + now[1]
    time_ranges_sec = []
    for t in time_ranges:
        time_range = sorted([x for x in t.values()])
        time_range = [int(t[0])*60+int(t[1]) for t in time_range]
        time_ranges_sec.append(time_range)
    for trs in time_ranges_sec:
        if trs[0]<=now_sec<=trs[1]:
            return 1
    return 0


def main(dp_endpoint, pipe_num, worker_endpoint, worker_start_no, worker_num, need_do_transform = False):
    all_worker_endpoint = []
    for i in xrange(worker_num):
        worker_port = worker_start_no + i
        all_worker_endpoint.append(worker_endpoint + str(worker_port))
    queues = []
    process_manger = []

    for i in range(5):
        q = Queue()
        msg_switch_handle = MsgSwitch(need_do_transform, dp_endpoint, all_worker_endpoint, pipe_num,q,True)
        process_manger.append(msg_switch_handle)
        queues.append(q)

    msg_switch_get = MsgSwitch(need_do_transform, dp_endpoint, all_worker_endpoint, pipe_num,queues,False)
    process_manger.append(msg_switch_get)

    for i in process_manger:
        i.start()
    for i in process_manger:
        i.join()


weekday = ['week_mon', 'week_tue', 'week_wed', 'week_thu', 'week_fri', 'week_sat', 'week_sun']
WEEKDAY = dict(zip(weekday, range(0,7)))


class MsgSwitch(Process):
    SendNOTSYNC=0
    SendNOTSYNCCount = 0
    def __init__(self,need_do_transform, dp_endpoint, all_worker_endpoint, pipe_num,queues,connected):
        super(MsgSwitch,self).__init__()
        self.instance_coll = get_instance_id_collection()
        self.dp_endpoint = dp_endpoint
        self.all_senders = []
        self.connected = connected
        self.file_path = os.path.join(settings.TMP_ROOT, 'file_switch')
        self.pipe_num = pipe_num
        self.msg_cache_map = {}
        self.msg_cache_prompt = 0
        self.msg_cache_map_count = {}
        self.app_intf_topov_is_twoway_trade = {}
        self.app_reverse_intf = {}
        self.app_intf_server_ips = {}
        self.msg_cache_total = 0
        self.last_ts = {}
        self.total_drop = 0
        self.btr_writer = BtrWriter()
        self.RedisBase =redis_operation.Redis()
        self.instances = None
        self.protocols = {}
        self.file_switch_last_modify = None
        self.file_switch = None
        self.need_do_transform=need_do_transform
        self.queues = queues
        self.all_worker_endpoint = all_worker_endpoint
        #
        # try:
        #     logger.debug('---start init---')
        #     self.init()
        # except Exception as ex:
        #     self.instances = None
        #     logger.error(traceback.format_exc(ex))

    def _is_twoway_trade(self, instance_id):
        app_name = self.instances[instance_id].get('app_name')
        intf_name = self.instances[instance_id].get('intf_name')
        if instance_id not in self.app_intf_topov_is_twoway_trade:
            topov_datapath = get_cbms_new_app_datapath(app_name)
            topov_capture = topov_datapath.get_capture(intf_name)
            is_twoway_trade = topov_capture.is_twoway_trade
            self.app_intf_topov_is_twoway_trade[instance_id] = is_twoway_trade
            if not is_twoway_trade:
                return is_twoway_trade
            infos = topov_datapath.get_twoway_trade_captures(intf_name)
            master_capture = infos['master']
            for intf, topov_capture in infos['slaves'].iteritems():
                self.app_reverse_intf[instance_id] = {"intf_name": intf, "server_ips": topov_capture.server_ips}
            self.app_intf_server_ips[instance_id] = master_capture.server_ips
        else:
            is_twoway_trade = self.app_intf_topov_is_twoway_trade[instance_id]
        return is_twoway_trade

    def set_reverse_intf(self, instance_id, rec):
        if not self._is_twoway_trade(instance_id):
            rec[REVERSE_INTF] = self.instances[instance_id]["intf_name"]
            return
        if rec.get(BTRKEY_IP_DST) in self.app_reverse_intf[instance_id]["server_ips"]:
            rec[REVERSE_INTF] = self.app_reverse_intf[instance_id]["intf_name"]
        else:
            rec[REVERSE_INTF] = self.instances[instance_id]["intf_name"]

    def get_transaction_key(self,instance_id ,rec):
        transaction_id = rec.get(DKEY_BIZ_TRANSACTION_ID, None)
        if not transaction_id:
            return None
        if self._is_twoway_trade(instance_id):
            if self.protocols[instance_id].should_start_transaction(rec):
                if rec.get(BTRKEY_IP_DST) in self.app_intf_server_ips[instance_id]:
                    return (instance_id, transaction_id, 'M')
                else:
                    return (instance_id, transaction_id, 'S')
            else:
                if rec.get(BTRKEY_IP_SRC) in self.app_intf_server_ips[instance_id]:
                    return (instance_id, transaction_id, 'M')
                else:
                    return (instance_id, transaction_id, 'S')
        else:
            return (instance_id, transaction_id)

    def init(self):
        if os.path.exists(self.file_path):
            self.file_switch_last_modify = os.stat(self.file_path).st_mtime
            with open(self.file_path, 'rb') as f:
                self.file_switch= True if f.read() == 'on' else False
        else:
            self.file_switch =None
        instances = list(self.instance_coll.find({}, {'_id': 0}))
        self.instances = {}
        for instance in instances:
            self.instances[instance.get('instance_id')] = instance
            hash_key = redis_operation.HASH_APP_INSTANCES % instance.get('app_name')
            hash_field=redis_operation.HASH_APP_INSTANCES_FIELD %instance.get('instance_id')
            self.RedisBase.save_hash(hash_key,hash_field,instance)
        self.transformers_cache={}
        self._get_transformers()
        self._get_protocol()
        if self.connected:
            self.sender_num = len(self.all_worker_endpoint)
            for worker_endpoint in self.all_worker_endpoint:
                self.all_senders.append(MsgSender(self.context, worker_endpoint))
            self.msg_send = MsgSender(self.context, "tcp://127.0.0.1:3333")

    def _get_transformers(self,instance_id=None):
        if not instance_id:
            for instance_id,instance in self.instances.iteritems():
                transformer=lookup_transformer(instance_id)
                self.transformers_cache[instance_id]=transformer
        else:
            transformer = lookup_transformer(instance_id)
            self.transformers_cache[instance_id] = transformer
            return transformer

    def _get_protocol(self,instance_id=None):
        if not instance_id:
            for instance_id,instance in self.instances.iteritems():
                protocol=lookup_protocol(instance_id)
                self.protocols[instance_id] = protocol
        else:
            protocol = lookup_protocol(instance_id)
            self.protocols[instance_id] = protocol

    def delete_cache_by_instance(self,instance_id):
        if instance_id in self.transformers_cache:
            del self.transformers_cache[instance_id]

    def transformers(self,instance_id):
        if not instance_id in self.transformers_cache:
            transformer = self._get_transformers(instance_id)
            return transformer
        else:
            return self.transformers_cache[instance_id]

    def run(self):
        self.context = zmq.Context()
        if not self.connected:
            try:
                input_channel = ZMQInputChannel(self.context, self.dp_endpoint, multipart=True, hwm=ZEROMQ_DEFAULT_HWM)
            except zmq.ZMQError as e:
                tc_logger.error('type=init tags="BACKOFF" message="ntr_switch, endpoint: %s, %s."' % (self.dp_endpoint, str(e)))
                raise
            if self.need_do_transform:
                i = 0
                records = iter(input_channel)
                for record in records:
                    try:
                        # if not self.instances:
                        #     self.init()
                        # t1 = time.time()
                        # gevent.spawn(self.transforme_handle,record)
                        # gevent.sleep(0)
                        ts1=time.time()
                        self.queues[i].put(record)
                        i+=1
                        i = i %len(self.queues)
                        # self.transforme_handle(record)
                        # t2=time.time()
                        # if t2 - t1 >= 0.05:
                        #     logger.info('-----find delay=%s in transforme_handle-----' % (t2 - t1))
                    except Exception as ex:
                        logger.error(traceback.format_exc(ex))
            else:
                records = input_channel
                for head_msg, body_msg in records:
                    head_info, body = decode_input((head_msg, body_msg))
                    bl=self.check_instance_id(head_info[2])
                    if not bl:
                        continue
                    inputs = [head_info, body]
                    self.handle(*inputs)
        else:
            while True:
                try:
                    if not self.instances:
                        self.init()
                    if self.queues.qsize() >100:
                        logger.debug("======queue===%s =================="% self.queues.qsize())
                    record = self.queues.get()
                    self.transforme_handle(record)
                except:
                    logger.error(traceback.format_exc())

    def check_instance_id(self,instance_id):
        instance=self.instances.get(instance_id)
        if instance:
            return True
        #     app_name,intf_name=instance.get('app_name'),instance.get('intf_name')
        #     hash_key=redis_operation.HASH_APP_INTFS_KEY %app_name
        #     hash_field=redis_operation.HASH_APP_INTFS_FIELD %intf_name
        #     bl=self.RedisBase.hash_field_exist(hash_key,hash_field)
        #     if bl:
        #         return True
        #     logger.error('----instance=%s not found in redis' %instance_id)
        # del self.instances[instance_id]
        self.delete_cache_by_instance(instance_id)
        return False

    def transforme_handle(self,record):
        # t0=time.time()
        head_info, body = decode_input(record)
        if not "DP#SYNC" in body:
            head_info[2] = body[DECODE_ID]
            ts = head_info[0]
            if time.time() - ts >10:
                logger.debug("===package ts   %s=============="%(time.time() - ts))
        # t1 = time.time()
        # if t1 - t0 >= 0.05:
        #     logger.info('-----find delay=%s in decode_input-----' % (t1 - t0))
        # ts = head_info[0]
        # if t1 - ts >= 20:
        #     logger.info('Msg delay=%s received with header:%s' % ((t1 - ts), head_info))
        instance_id = head_info[2]
        # t2 = time.time()
        if head_info[4] != 0:
            transformer = self.transformers(instance_id)
            body = transformer(body)
        head_info, body=make_hash_id(head_info, body,self)
        # t3 = time.time()
        # if t3 - t2 >= 0.1:
        #     logger.info('-----find delay=%s in transformer-----' % (t3 - t2))
        # inputs =[head_info, body]
        return
        self.send_msgs(head_info,body)

    def send_msgs(self,head_info,body_info):
        if head_info[4] == 0:
            sync_info = {BTRKEY_TS: head_info[0] * 1000000000,
                         BTRKEY_DP_SYNC: 1}
            sync_body_msg = pack_msg(sync_info)
            sync_head_msg = pack_msg([head_info[0],
                                      0,
                                      head_info[2],
                                      0,
                                      0,
                                      0])
            # gevent.spawn(self.all_senders[head_info[0] % self.sender_num].send,head_info[0], sync_head_msg, sync_body_msg, True)
            self.all_senders[head_info[0] % self.sender_num].send( head_info[0], sync_head_msg, sync_body_msg, True)
        else:
            # gevent.spawn(self.all_senders[head_info[4] % self.sender_num].send,head_info[0], pack_msg(head_info),pack_msg(body_info),False)
            self.all_senders[head_info[4] % self.sender_num].send( head_info[0], pack_msg(head_info), pack_msg(
                body_info), False)
    @property
    def judge_file_switch(self):
        # t0=time.time()
        mtime= os.stat(self.file_path).st_mtime
        # t1=time.time()
        # if t1-t0>=0.05:
        #     logger.info('---find delay=%s in judge_file_switch---' %(t1-t0))
        if self.file_switch_last_modify != mtime:
            self.file_switch_last_modify = mtime
            self.file_switch = not self.file_switch
            return self.file_switch
        return self.file_switch

    def handle(self,head_info,body):
        # t0=time.time()
        instance_id = head_info[2]
        ts = head_info[0]
        instance=self.instances.get(instance_id)
        worktime = instance.get('worktime', 0)
        switch = instance.get('worktime_swicth', 'on')
        msg_ts = time.localtime(ts)
        now = [msg_ts[3], msg_ts[4]]
        if switch == 'on' and worktime:
            access_flag = 0
            for wk in worktime:
                if msg_ts[6] in [WEEKDAY[w] for w in wk['week']]:
                    if cmp_time(now, wk['time_range']):
                        access_flag = 1
                        break
            if not access_flag:
                head_info[4] = 0

        if self.judge_file_switch:
            headfile =head_info
            bodyfile = body
            if 'DP#SYNC' not in bodyfile:
                suffix = time.strftime("%Y_%m_%d", time.localtime())
                BASE_DIR = os.path.abspath(settings.TMP_ROOT)
                dirh = os.path.join(BASE_DIR, 'headfile_'+suffix)
                dirb = os.path.join(BASE_DIR, 'bodyfile_'+suffix)
                with open(dirh,"a+") as f:
                    f.write(str(headfile)+'\n')
                with open(dirb,"a+") as f:
                    f.write(str(bodyfile)+'\n')
        self.switch(head_info,body)
        # t1 = time.time()
        # if t1 - t0 >= 0.1:
        #     logger.info('-----find delay=%s in handle-----' % (t1 - t0))

    def switch(self,head_info, body):
        # t1 = time.time()
        try:
            ts_sec, ts_nsec, instance_id, pipe_id, hash_key, reserved = head_info
            last_ts_key = (instance_id, pipe_id)
            try:
                last_ts_val = self.last_ts[last_ts_key]
            except KeyError:
                last_ts_val = [0, 0]
                self.last_ts[last_ts_key] = last_ts_val

            if ts_sec < last_ts_val[0]:
                last_ts_val[1] += 1
                self.total_drop += 1
                return
            if ts_sec > last_ts_val[0]:
                last_ts_val[0] = ts_sec
            try:
                msg_cache = self.msg_cache_map[instance_id]
            except KeyError:
                msg_cache = MsgCache(instance_id, self.pipe_num, self.sender_num)
                self.msg_cache_map[instance_id] = msg_cache
            head_msg, body_msg = [pack_msg(head_info), pack_msg(body)]
            inputs=[ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg]
            # ti=time.time()
            to_pop_msgs = msg_cache.push_pop_msg(*inputs)
            # t2 = time.time()
            # if t2 - t1 >= 0.1:
            #     logger.info('----find delay=%s in handle_2-----' % (t2 - t1))
            #     if t2-ti>=0.1:
            #         logger.info('----find delay ti=%s-----' % (t2 - ti))
            self.send_msg(to_pop_msgs)
        except Exception as ex:
            logger.error(traceback.format_exc(ex))

    def send_msg(self,to_pop_msgs):
        # t2 = time.time()
        for sender_no, ts, head_msg, body_msg, is_sync in to_pop_msgs:
            self.all_senders[sender_no].send(ts, head_msg, body_msg, is_sync)
        # t3 = time.time()
        # if t3 - t2 >= 0.1:
        #     logger.info('------find delay=%s in send_msg-----' % (t3 - t2))

    def insert_to_redis(self,head_info, body):
        if "DP#SYNC" not in body:
            instance_id = head_info[2]
            instance= self.instances.get(instance_id)
            app_name, intf_name =(instance.get('app_name'),instance.get('intf_name'))
            key = redis_operation.RAW_DATA_KEY % (app_name, intf_name)
            body["app_name"] = app_name
            body["intf_name"] = intf_name
            body['ts'] = body['ts'] * 1.0 / 1000000000
            # t0=time.time()
            bl=self.RedisBase.insert_to_list(key,body)
            # t1 = time.time()
            # if t1 - t0 >= 0.2:
            #     logger.info('---find delay=%s in insert_to_redis---' % (t1 - t0))
            if not bl:
                logger.error('---insert_to_redis fail !---')

class MsgCache(object):
    """
        cache messages for each instance_id.
        ensure the pop messages is time sorted for each instance_id, and send sync messages
    """
    NOThash_key = 0
    def __init__(self, instance_id, pipe_num, sender_num):
        self.instance_id = instance_id
        self.sender_num = sender_num
        self.sender_no_list = range(sender_num)
        self.sync_ts = None
        self.pipe_num = pipe_num
        self.pipe_msg_cnt = [0] * pipe_num
        self.pipe_cache_msgs = []
        for i in xrange(pipe_num):
            self.pipe_cache_msgs.append(deque())
        self.ts=time.time()
        self.sync_time=time.time()
        self.cache_msgs=[]

    def old_push_pop_msg(self, ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg):
        self.push_msg(ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg)
        return self.pop_msgs()

    def push_msg(self, ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg):
        self.pipe_cache_msgs[pipe_id].append((ts_sec,
         ts_nsec,
         pipe_id,
         hash_key,
         head_msg,
         body_msg))
        self.pipe_msg_cnt[pipe_id] += 1

    def  push_pop_msg(self, ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg):
        self.cache_msgs.append((ts_sec,
         ts_nsec,
         hash_key,
         head_msg,
         body_msg))
        to_pop_msgs=[]
        for msg in self.cache_msgs:
            res=self.gen_msgs(msg)
            if res:
                to_pop_msgs+=res
        self.cache_msgs=[]
        return to_pop_msgs

    def pop_msgs(self, sync_delta = 1):
        to_pop_msgs = []
        enter=False
        while max(self.pipe_msg_cnt)>0:
            # if not enter:
                # now=time.time()
                # if now-self.ts>=0.1:
                #     logger.info('------pop_msgs with timespan:%s---------' %(now-self.ts))
            self.ts=time.time()
            enter=True
            info=[]
            for i,msgs in enumerate(self.pipe_cache_msgs):
                if self.pipe_msg_cnt[i]>0:
                    info.append(msgs[0])
            ts_sec, ts_nsec, pipe_id, hash_key, head_msg, body_msg=min(info)
            self.pipe_cache_msgs[pipe_id].popleft()
            self.pipe_msg_cnt[pipe_id] -= 1
            try:
                to_sync_ts = ts_sec // sync_delta * sync_delta
                if to_sync_ts - self.sync_ts > 0:
                    tnow=time.time()
                    # if tnow-self.sync_time>=0.1:
                    #     logger.info('---------to send with timespan:%s----------' %(tnow-self.sync_time))
                    self.sync_time=tnow
                    sync_head_msg, sync_body_msg = self.gen_sync_msg(to_sync_ts)
                    for sender_no in self.sender_no_list:
                        to_pop_msgs.append((sender_no,
                         to_sync_ts,
                         sync_head_msg,
                         sync_body_msg,
                         True))

                    self.sync_ts = to_sync_ts
            except TypeError:
                self.sync_ts = to_sync_ts

            if hash_key:
                sender_no = hash_key % self.sender_num
                to_pop_msgs.append((sender_no,
                 ts_sec,
                 head_msg,
                 body_msg,
                 False))

        return to_pop_msgs

    def gen_msgs(self,msg,sync_delta=1,timeout=1):
        res=[]
        ts_sec, ts_nsec, hash_key, head_msg, body_msg = msg
        if hash_key:
            sender_no = hash_key % self.sender_num
        else:
            sender_no=None
        to_sync_ts = ts_sec // sync_delta * sync_delta
        if not self.sync_ts:
            self.sync_ts=to_sync_ts
        elif to_sync_ts - self.sync_ts > 0 or  time.time()-self.sync_time>timeout:
            tnow = time.time()
            # if tnow - self.sync_time >= timeout:
            #     logger.info('---------to send with timespan:%s----------' % (tnow - self.sync_time))
            self.sync_time = tnow
            self.sync_ts = to_sync_ts
            sync_head_msg, sync_body_msg = self.gen_sync_msg(to_sync_ts)
            if sender_no:
                res.append((sender_no,
                            to_sync_ts,
                            sync_head_msg,
                            sync_body_msg,
                            True))
            else:
                for sender_no in self.sender_no_list:
                    res.append((sender_no,
                                        to_sync_ts,
                                        sync_head_msg,
                                        sync_body_msg,
                                        True))
        if hash_key:
            res.append((sender_no,
                                ts_sec,
                                head_msg,
                                body_msg,
                                False))
        return res

    def gen_sync_msg(self, sync_ts):
        sync_info = {BTRKEY_TS: sync_ts * 1000000000,
         BTRKEY_DP_SYNC: 1}
        sync_body_msg = pack_msg(sync_info)
        sync_head_msg = pack_msg([sync_ts,
         0,
         self.instance_id,
         0,
         0,
         0])
        return (sync_head_msg, sync_body_msg)


class MsgSender(object):
    """
        send message format is:
        [
            head_msg: [
                TS_SEC, TS_NSEC, INSTANCE_ID, PIPE_ID, HASH_KEY, RESERVED
            ],
            body_msg: MESSAGE_RECORD | SYNC_RECORD
        ]
    """

    def __init__(self, context, send_endpoint):
        sender = context.socket(zmq.PUSH)
        sender.setsockopt(zmq.SNDHWM, ZEROMQ_DEFAULT_HWM)
        sender.connect(send_endpoint)
        self.send_multipart = sender.send_multipart
        self.send_endpoint = send_endpoint
        self.msg_sent_cnt = 0
        self.msg_buffer = []
        self.ts=time.time()

    def send(self, ts, header_msg, body_msg, is_sync = False):
        if is_sync:
            now=time.time()
            self.ts=now
            if self.msg_buffer:
                for header_msg_b, body_msg_b in self.msg_buffer:
                    self.send_multipart([header_msg_b, body_msg_b])
                    self.msg_sent_cnt += 1
            self.send_multipart([header_msg, body_msg])

            self.msg_sent_cnt += 1
            self.msg_buffer = []
            if self.msg_sent_cnt >= 10000:
                logger.debug('Sender[%s]: send last message %s' % (self.send_endpoint, ts))
                self.msg_sent_cnt = 0
        else:
            self.msg_buffer.append((header_msg, body_msg))
            pass
        # gevent.sleep(0)

class BtrWriter(object):

    def __init__(self):
        self._store_root = getattr(settings, 'STORE_ROOT')
        self._caches = {}
        self.instance_coll = get_instance_id_collection()

    def is_double_center(self, instance_id):
        if instance_id in self._caches:
            if self._caches[instance_id]['is_double_live']:
                return True
            else:
                return False
        else:
            instance = self.instance_coll.find_one({'instance_id': instance_id})
            if instance:
                topov = get_cbms_new_app_datapath(instance['app_name'])
                capture = topov.get_capture(instance['intf_name'])
                if capture.is_double_live:
                    self._caches[instance_id] = {'part': '%(app_name)s/%(intf_name)s' % {'app_name': instance['app_name'],
                              'intf_name': instance['intf_name']},
                     'is_double_live': True}
                    return True
                self._caches[instance_id] = {'is_double_live': False}
        return False

    def _get_btr_file(self, ts, instance_id, pipe_id):
        # part_path = self._caches[instance_id]['part']
        current_minute = ts // 60 * 60
        formated_datetime = datetime.datetime.fromtimestamp(current_minute).strftime('%Y%m%d%H%M%S')
        btr_filename = '%(ft)s_%(instance)s_%(pipe)s.btr.pack' % {'ft': formated_datetime,'instance':instance_id,'pipe': str(pipe_id)}
        # btr_fullpath = os.path.join(self._store_root, 'btr', part_path, btr_filename)
        btr_fullpath = os.path.join(self._store_root, 'btr', btr_filename)
        return btr_fullpath

    def write(self, msgs):
        to_send_msgs = [ (unpack_msg(head), body) for head, body in msgs] # if self.is_double_center(unpack_msg(head)[2]) ]
        for head, body in to_send_msgs:
            btr_file = self._get_btr_file(head[0], head[2], head[3])
            with open(btr_file, 'a') as fp:
                fp.write(body)

def find_instance(RedisBase,instance_id):
    hash_key_pattern=redis_operation.HASH_APP_INSTANCES_PATTERN
    keys=RedisBase.find_keys(hash_key_pattern)
    hash_field=redis_operation.HASH_APP_INSTANCES_FIELD % instance_id
    instance=None
    for hash_key in keys:
        instance=RedisBase.query_hash(hash_key,hash_field,query_mongo=True)
        if instance:
            return instance

def get_all_instances(RedisBase):
    result={}
    hash_key_pattern = redis_operation.HASH_APP_INSTANCES_PATTERN
    keys = RedisBase.find_keys(hash_key_pattern)
    for hash_key in keys:
        data=RedisBase.query_hash(hash_key)
        result.update(data)
    return result

if __name__ == '__main__':
    if len(sys.argv) >= 6:
        dp_endpoint = sys.argv[1]
        pipe_num = int(sys.argv[2])
        worker_endpoint = sys.argv[3]
        worker_start_no = int(sys.argv[4])
        worker_num = int(sys.argv[5])
        need_do_transform = len(sys.argv) >= 7 and sys.argv[6] == 'btr'
        main(dp_endpoint, pipe_num, worker_endpoint, worker_start_no, worker_num, need_do_transform)
    else:
        print 'Usage: %s dp_endpoint pipe_num worker_endpoint worker_start_no worker_num' % sys.argv[0]
        sys.exit(2)
