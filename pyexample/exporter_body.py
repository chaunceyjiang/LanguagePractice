'''
{'ip_src': '11.194.16.51',
 'ip_dst': '11.194.32.33',
 'ret_code': None,
 'sys_type': 'IBPS',
 'reverse_intf': 'intf10',
 'CnccMesgType': 'ccms.990.001.01     ',
 'sub_trans_type': 'ccms',
 'trans_type': '\xe9\x80\x9a\xe4\xbf\xa1\xe7\xba\xa7\xe7\xa1\xae\xe8\xae\xa4',
 'sys_status': None,
 'org_status': None,
 'ts': 1539051960,
 'trans_count': 1.0,
 'succ_count': 1.0,
 'req_count': 1.0,
 'resp_count': 1.0,
 'duration': 0.0,
 'succ_rate': 100.0,
 'resp_rate': 100.0,
 'ts_end': 1539051960}
'''

'''
{'ip_src': '11.194.16.51',
 'ip_dst': '11.194.32.33',
 'ret_code': None,
 'sys_type': 'IBPS',
 'reverse_intf': 'intf10',
 'CnccMesgType': 'ccms.990.001.01     ',
 'sub_trans_type': 'ccms',
 'trans_type': '\xe9\x80\x9a\xe4\xbf\xa1\xe7\xba\xa7\xe7\xa1\xae\xe8\xae\xa4',
 'sys_status': None,
 'org_status': None,
 'ts': 1539051960,
 'trans_count': 1.0,
 'succ_count': 1.0,
 'req_count': 1.0,
 'resp_count': 1.0,
 'duration': 0.0,
 'succ_rate': 100.0,
 'resp_rate': 100.0,
 'ts_end': 1539051960}
'''
import os,sys
PATH="/home/chauncey/Desktop/exporter_bodyfile_2018_10_09"
PATH_1="/home/chauncey/Desktop/exporter_bodyfile_2018_10_09_1"
def reqSum(path):
    with open(path,'r') as f:
        body = f.readlines()
    req_sum = 0
    for i in body:
        req = int(eval(i).get("req_count"))
        req_sum += req
    return req_sum

def getTs(path):
    ts_list =[]
    with open(path,'r') as f:
        body = f.readlines()
    for i in body:
        ts = int(eval(i).get("ts_end"))
        ts_list.append(ts)
    print len(ts_list)

# getTs(PATH_1)
# print reqSum(sys.argv[1])
# print reqSum(PATH_1)
# print reqSum(PATH)
'''
38339
47022
'''
# print reqSum(PATH_1)

def Odd_number():
    i = 2
    while i:
        i+=2
        yield i

def Even_number():
    i = 1
    while i:
        i+=2
        yield i
import time
f= Odd_number()
ff =Even_number()
while True:
    print ff.next()
    print f.next()

    print "============"
    time.sleep(1)