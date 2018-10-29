# coding: utf-8
import msgpack
import sys
import os.path
_path = os.path.abspath(os.path.join(os.path.dirname(__file__), '../../BMS/cbms'))

# _path2 = os.path.abspath(os.path.join(os.path.dirname(__file__), './BMS/cbms/')
if _path not in sys.path:
    sys.path.append(_path)
    # sys.path.append(_path2)
from django.conf import settings
if not settings.configured:
    os.environ['DJANGO_SETTINGS_MODULE'] = 'settings'
from dataprovider.transform import ProtocolTransformPipe
from dataprovider.protocol.config import make_protocol_by_name
from dataprovider.transaction import make_transaction_pipe
def unpack_raw(stream):
    unpacker = msgpack.Unpacker(stream)
    for msg in unpacker:
        yield msg

protocol  = make_protocol_by_name("cncchvps")

def unpack_msg(msg):
    return msgpack.unpackb(msg)


def pack_msg(pyobj):
    return msgpack.packb(pyobj)

PATH = "/home/chauncey/Desktop/tmp/"
def decode():
    f = open(PATH+"20181022152300.btr.sync","a+")

    for i in range(0,1):
        filename = PATH+"20181022152300_20181022152300_"+str(i)+ ".btr.pack"
        print filename
        records = ProtocolTransformPipe(protocol)(unpack_raw(file(filename)))
        for record in records:
            # if 'DP#SYNC' not in record:
            f.write(str(record)+"\n")
    f.close()
decode()
def count_ret_code():
    PR04 = 0
    null = 0
    ff = open(PATH+"20181022152200.btr","r")
    records = []
    for record in ff.readlines():
        records.append(eval(record))
    for record in records:
        if record.get("RRA") == "resp":
            if  record.get("ret_code") == "PR04":
                PR04 +=1
            else:
                null +=1 
    print PR04,null

# count_ret_code()


def count_transaction_id():
    PR04 = 0
    dest_ip_list = [
    "11.194.16.16",
    "11.194.16.17",
    "11.194.16.18",
    "11.194.16.19",
    ]
    null = 0
    req = 0
    req_2 = 0
    S = {}
    ff = open(PATH+"20181022152200.btr","r")
    f = open(PATH + "20181022152300.btr","r")
    records = []
    for record in ff.readlines():
        records.append(eval(record))
    # for record in f.readlines():
    #     records.append(eval(record))

    for record in records:
        ts = (record.get('ts') / 1000000000 ) // 60 * 60

        # if record.get("RRA") == "resp":
        # if record.get("RRA") == "req" or record.get("RRA") == "notify":
        if record.get("RRA") == "req" or record.get("RRA") == "notify": 
            if record.get("transaction_id") not in S and record.get("DestIp") in dest_ip_list:
                if ts < 1540192980:
                    S[record.get("transaction_id")] = [record]

        if record.get("RRA") == "resp":
            if record.get("transaction_id") in S and record.get("SrcIp") in dest_ip_list:
                S[record.get("transaction_id")].append(record)

    for record in records:
        ts = (record.get('ts') / 1000000000 ) // 60 * 60
        
        # if record.get("RRA") == "resp":
        # if record.get("RRA") == "req" or record.get("RRA") == "notify":
        if record.get("RRA") == "req" or record.get("RRA") == "notify": 
            if record.get("transaction_id") not in S and record.get("DestIp") in dest_ip_list:
                if ts < 1540192980:
                    S[record.get("transaction_id")] = [record]

        if record.get("RRA") == "resp":
            if record.get("transaction_id") in S and record.get("SrcIp") in dest_ip_list:
                S[record.get("transaction_id")].append(record)

            # else:
            #     print "====================================="
            #     print record
            #     print S[record.get("transaction_id")] 
            #     print "====================================="
            #     req += 1

        # if  record.get("transaction_id") not in S:
        #     S[record.get("transaction_id")] = record.get("transaction_id")
        # else:
        #     print record
    for transaction_id in S:
        if len(S[transaction_id]) >=2:
            req_2 += 1
            print S[transaction_id]
    print req_2,len(S)

# count_transaction_id()





# def make_transaction():
#     tran_cache = {}
#     _build = 0
#     _build_invalid = 0
#     f = open(PATH+"20181022152200.btr","r")
#     records = f.readlines()
#     for record in records:
#         record = eval(record)
#         transaction_id = record.get('transaction_id')
#         if transaction_id not in tran_cache:
#             tran_cache[transaction_id] = record
#         else:
#             print  tran_cache[transaction_id],record
#             if tran_cache[transaction_id].get("RRA") == "req" and record.get("RRA")=="resp" :
#                 _build += 1
#             elif tran_cache[transaction_id].get("RRA") == "resp" and record.get("RRA")=="req":
#                 _build += 1
#             del tran_cache[transaction_id]
#         _build_invalid = len(tran_cache)
#     print ""
#     print _build,_build_invalid
# transaction_pipe = make_transaction_pipe(protocol)
# def make_transaction():
#     f = open(PATH+"20181022152200.btr","r")
#     ff = open(PATH+"20181022152200.btr.tran","a+")
#     records = f.readlines()
#     records_eval = []
#     for i in records:
#         records_eval.append(eval(i))
#     transactions =  transaction_pipe(records_eval)
#     for transaction in transactions:
#         # print transaction
#         ff.write(str(transaction)+"\n")
# make_transaction()

# def count_req():
#     ff = open(PATH+"20181022152200.btr.tran","r")
#     records = []
#     for record in ff.readlines():
#         records.append(eval(record))

















