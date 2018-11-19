import time
import logging
from collections import defaultdict
from threading import Thread
# from threading import Timer
from confintf.data.btr_key import BTRKEY_DP_SYNC
from confintf.mongo import get_duration_peak_collection
from confintf.data.db import DKEY_TS, DKEY_BIZ_DURATION
from confintf.datapath import get_cbms_new_app_datapath
from django.conf import settings
import os
# from dataexporter.models import RecordNormalizer
# from datastore import redis_operation
# RedisBase = redis_operation.Redis()
logger = logging.getLogger('default')

class DurationPeakPipe(object):
    stats_interval = 60
    db_coll = get_duration_peak_collection()
    _cache = {}

    def __init__(self, app_name, intf_name):
        self.app_name = app_name
        self.intf_name = intf_name
        self.max_insert_ts = 0
        self.max_value_cache = defaultdict(dict)
        self._path = os.path.join(settings.STORE_ROOT,'durationPeak')
        # self.normalizer = RecordNormalizer()

    def __call__(self, records):
        # twoway_trade_info = self.get_twoway_trade_info(self.app_name, self.intf_name)
        for rec in records:
            t0=time.time()
            reverse_intf = rec.get('reverse_intf')
            if reverse_intf:
                self.intf_name = reverse_intf
                # record = self.normalizer.normalize(self.app_name, self.intf_name, rec)
                # if not twoway_trade_info['twoway']:
                #     if record['dst_server']:
            self.process(rec)
                # elif 'intfs' in twoway_trade_info:
                #     twoway_intfs = filter(lambda x: x['intf_name'] != self.intf_name, twoway_trade_info['intfs'])
                #     reverse_intf = self.find_reverse_intf(record, twoway_intfs)
                #     if reverse_intf:
                #         self.intf_name = reverse_intf
                        # t0 = time.time()
                        # self.process(record)
                        # t1 = time.time()
                        # if t1 - t0 >= 0.1:
                        #     logger.info('-----find delay2=%s in process-----' % (t1 - t0))
                        # logger.debug('Drop bar in twoway transaction, for cannot find reverse direction intf')
            if rec.get('transaction_id') == "2018102600000213":
                logger.debug("===============DurationPeakPipe %s=========="%(self.max_value_cache))
            t1 = time.time()
            if t1 - t0 >= 0.1:
                logger.debug('-----find delay=%s in DurationPeakPipe-----' % (t1 - t0))
            yield rec

    def process(self, record):
        ts, max_value = self.prepare(None, record)
        # logger.debug("==========process ts  %s======cache %s ==========="%(ts,self.max_value_cache))
        if BTRKEY_DP_SYNC not in record:
            trans_type = record.get('trans_type')
            sub_trans_type = record.get('sub_trans_type')
            if ts > self.max_insert_ts:
                cache_value = self.max_value_cache[ts].get((self.app_name, record.get('reverse_intf'), trans_type, sub_trans_type))
                if cache_value is None:
                    if max_value:
                        self.max_value_cache[ts][(self.app_name, record.get('reverse_intf'), trans_type, sub_trans_type)] = max_value
                else:
                    if max_value > cache_value:
                        self.max_value_cache[ts][(self.app_name, record.get('reverse_intf'), trans_type, sub_trans_type)] = max_value
            else:
                if max_value:
                    _instance_path = os.path.join(self._path, self.app_name, record.get('reverse_intf'), str(ts))
                    # self.db_coll.insert({'ts': ts, 'ts_end': ts + 59, 'app_name': self.app_name,
                    #             'intf_name': self.intf_name,'trans_type': trans_type, 'sub_trans_type':
                    #             sub_trans_type, 'peak': max_value})
                    if not os.path.exists(_instance_path):
                        try:
                            os.makedirs(_instance_path)
                        except OSError:
                            pass
                    with open(_instance_path+os.sep+trans_type+"_"+sub_trans_type,"a+") as f:
                        f.write(str(max_value)+"\n")

        if self.max_value_cache:
            min_ts = min(self.max_value_cache.keys())
            if ts > min_ts:
                insert_value = [{'ts': min_ts, 'ts_end': min_ts + 59, 'app_name': _k[0], 'intf_name': _k[1],
                         'trans_type': _k[2], 'sub_trans_type': _k[3], 'peak': _v}
                                for _k, _v in self.max_value_cache[min_ts].iteritems()]

                for pre_insert_value in insert_value:
                    _instance_path = os.path.join(self._path, pre_insert_value['app_name'], pre_insert_value['intf_name'], str(pre_insert_value['ts']))
                    if not os.path.exists(_instance_path):
                        try:
                            os.makedirs(_instance_path)
                        except OSError:
                            pass
                    with open(_instance_path + os.sep + pre_insert_value['trans_type']+"_"+pre_insert_value['sub_trans_type'], "a+") as f:
                        f.write(str(pre_insert_value['peak'])+"\n")
                self.max_insert_ts = min_ts
                # if insert_value:
                #     Thread(target=self.db_coll.insert, args=(insert_value, )).start()
                del self.max_value_cache[min_ts]

        # logger.debug("==========process_2  ts  %s======cache %s ===========" % (ts, self.max_value_cache))

    def prepare(self, header_info, records):
        ts = records.get(DKEY_TS) // self.stats_interval * self.stats_interval
        max_value = records.get(DKEY_BIZ_DURATION,0)
        return (ts, max_value)

    def get_twoway_trade_info(self, app_name, intf_name):
        cache_key = (app_name, intf_name)
        if cache_key not in self._cache:
            self._cache[cache_key] = self._get_twoway_trade_info(app_name, intf_name)
        return self._cache[cache_key]

    def _get_twoway_trade_info(self, app_name, intf_name):
        topov_datapath = get_cbms_new_app_datapath(app_name)
        topov_capture = topov_datapath.get_capture(intf_name)
        is_twoway_trade = topov_capture.is_twoway_trade
        if is_twoway_trade:
            intfs_infos = []
            infos = topov_datapath.get_twoway_trade_captures(intf_name)
            master_capture = infos['master']
            intfs_infos.append({'intf_name': master_capture['name'],
             'server_ips': master_capture.server_ips})
            for intf, topov_capture in infos['slaves'].iteritems():
                intfs_infos.append({'intf_name': intf,
                 'server_ips': topov_capture.server_ips})

            return {'twoway': True,
             'intfs': intfs_infos}
        else:
            return {'twoway': False}

    def find_reverse_intf(self, record, twoway_intfs):
        server_ip = record['ip_dst']
        for item in twoway_intfs:
            if server_ip in item['server_ips']:
                return item['intf_name']

