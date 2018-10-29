# # coding: utf-8
# '''
# {'ts': 1538982853053500000,

#  'FlowId': '847257487030755360:-6814216840525381631:0',

#  'PktLen': 768,

#  'MetaType': 1,

#  'StreamId': 1,

#  'PartId': 0,

#  'DecodeId': -971229470,

#  'FlowSide': 0,

#  'SrcIp': '11.194.16.50',

#  'DestIp': '11.194.32.32',

#  'SrcPort': 41327,

#  'DestPort': 2531,

#  'IpProt': 6,

#  'TcpPldLen': 714,

#  'PktId': 705410,

#  'TcpSeq': 1050567590,

#  'TcpAck': 2847887183,

#  'CnccOrigReceiverSID': 'IBPS',

#  'CnccMesgType': 'ibps.101.001.02     ',

#  'CnccMesgID': '30000000002233925355',

#  'CnccMesgRefID': '14040668353949641447',

#  'MsgId': '1021000999962018052435394964',

#  'tranType': 'C200',

#  'RRA': 'req',

#  'Prot': 'xml',

#  'RRB': 'req',

#  'mesgtype': 'ibps.101',

#  'sub_trans_type': 'ibps',

#  'transaction_id': '1021000999962018052435394964',

#  'ccmstype': ('ccms.303',

#  'ccms.307',

#  'ccms.308',

#  'ccms.310',

#  'ccms.311',

#  'ccms.312',

#  'ccms.313',

#  'ccms.314',

#  'ccms.315',

#  'ccms.316',

#  'ccms.317',

#  'ccms.318',

#  'ccms.319',

#  'ccms.801',

#  'ccms.803',

#  'ccms.805',

#  'ccms.806',

#  'ccms.807',

#  'ccms.809',

#  'ccms.811',

#  'ccms.900',

#  'ccms.903',

#  'ccms.906',

#  'ccms.907',

#  'ccms.911',

#  'ccms.913',

#  'ccms.915',

#  'ccms.916',

#  'ccms.917',

#  'ccms.919',

#  'ccms.921',

#  'ccms.926',

#  'ccms.990',

#  'ccms.991',

#  'ccms.992'),

#  'ibpstype': ('ibps.311',

#  'ibps.309',

#  'ibps.305',

#  'ibps.331',

#  'ibps.333',

#  'ibps.335',

#  'ibps.337',

#  'ibps.339',

#  'ibps.341',

#  'ibps.343',

#  'ibps.703',

#  'ibps.705'),

#  'trans_type': '\xe6\xb1\x87\xe5\x85\x91',

#  'ret_code': None,

#  'status': None,

#  'sys_type': 'IBPS',

#  'ms_key': [-971229470,

#  '1021000999962018052435394964',

#  'S'],

#  'reverse_intf': u'intf10'}



# '''
# # with open('/home/chauncey/Desktop/bodyfile_2018_10_08',

# #     body_switch = f.readlines()

# #     body_worker = f.readlines()
# # ~/Desktop/tmp/bodyfile_2018_10_10_4
# with open('/home/chauncey/Desktop/tmp/20181017103700_20181017103700_0.btr','r') as f:
#     body_worker = f.readlines()


# # ms_key_list_switch = []
# # for recodr in body_switch:
# #     key = eval(recodr).get('ms_key')
# #     try:
# #         key = tuple(key)
# #     except TypeError:
# #         print recodr
# #     ms_key_list_switch.append(key)

# print "-------------------------------------------------------------------"

# # ms_key_list_worker = []
# # for recodr in body_worker:
# #     key = eval(recodr).get('ms_key')
# #     try:
# #         key = tuple(key)
# #     except TypeError:
# #         print recodr
# #     ms_key_list_worker.append(key)
# # print ms_key_list_worker[1:10]
# # print len(set(ms_key_list_worker))
# # for i in ms_key_list_switch:
# #     if i in ms_key_list_worker:
# #         index = ms_key_list_worker.index(i)
# #         ms_key_list_worker.pop(index)

# # for i in ms_key_list_worker:
# #     if i in ms_key_list_switch:
# #         index = ms_key_list_switch.index(i)
# #         ms_key_list_switch.pop(index)
# # print ms_key_list_switch

# # reqSUM = 0
# # SET = set()
# # for i in body_worker:
# #     recodr = eval(i)
# #     key = tuple(recodr.get('ms_key'))
# #     if key in SET:
# #         print recodr
# #     if recodr.get("RRA") == "notify":
# #         reqSUM+=int(recodr.get('NbOfTxs',
# # 1))
# #         if key not in SET:
# #             SET.add(key)
# #         else:
# #             print "---------------"
# #             print recodr
# #             print "---------------"

# # print reqSUM,
# len(SET)

# # 47038
# # 47038 39333
# #13696 13695 notify 的个数，有一个重复的ms_key
# # {'CnccOrigReceiverSID': 'IBPS',
#  'status': None,
#  'ms_key': [-971229470,
#  '20180524100118029924',
#  'S'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'CnccMesgRefID': '20180524100118029924',
#  'Prot': 'xml',
#  'PktLen': 1111,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'notify',
#  'FlowId': '847258242944999456:-5232327471411494911:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 541414,
#  'CnccMesgID': '20180524100118029924',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.226',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
# 'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'IBPS',
#  'TcpPldLen': 1057,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.990',
#  'ret_code': None,
#  'SrcPort': 46947,
#  'ts': 1538982641560452000,
#  'RRA': 'notify',
#  'transaction_id': '20180524100118029924',
#  'IpProt': 6,
#  'trans_type': '\xe9\x80\x9a\xe4\xbf\xa1\xe7\xba\xa7\xe7\xa1\xae\xe8\xae\xa4',
#  'TcpSeq': 430641618,
#  'TcpAck': 349912811}


# # {'ts': 1538982605056191000,
#  'FlowId': '847258255829901347:-9047157830771015679:0',
#  'PktLen': 1123,
#  'MetaType': 1,
#  'StreamId': 1,
#  'PartId': 1,
#  'DecodeId': -971229470,
#  'FlowSide': 0,
#  'SrcIp': '11.194.16.229',
#  'DestIp': '11.194.32.35',
#  'SrcPort': 33394,
#  'DestPort': 2531,
#  'IpProt': 6,
#  'TcpPldLen': 1057,
#  'PktId': 520135,
#  'TcpSeq': 1617354085,
#  'TcpAck': 4137591923,
#  'CnccOrigReceiverSID': 'IBPS',
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'CnccMesgID': '20180524100118029924',
#  'CnccMesgRefID': '20180524100118029924',
#  'RRA': 'notify',
#  'Prot': 'xml',
#  'RRB': 'notify',
#  'mesgtype': 'ccms.990',
#  'sub_trans_type': 'ccms',
#  'transaction_id': '20180524100118029924',
#  'ccmstype': ('ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'),
#  'ibpstype': ('ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'),
#  'trans_type': '\xe9\x80\x9a\xe4\xbf\xa1\xe7\xba\xa7\xe7\xa1\xae\xe8\xae\xa4',
#  'ret_code': None,
#  'status': None,
#  'sys_type': 'IBPS',
#  'ms_key': [-971229470,
#  '20180524100118029924',
#  'S'],
#  'reverse_intf': u'intf10'}


# #




# #--------------------

# """  
# [{'HeaderType': 'U',
#  'Prot': 'cncc 2g',
#  'PktLen': 1191,
#  'ts': 1538982486492833000,
#  'PktId': 367791,
#  'SrcIp': '11.196.163.1',
#  'FlowSide': 0,
#  'mbfe_sts': '10:34:37.647591',
#  'TcpPldLen': 1137,
#  'ms_key': [-966899370,
#  '2018052481056594',
#  'M'],
#  'SrcPort': 61514,
#  'RRB': 'req',
#  'RRA': 'req',
#  'ccpc_cts': '10:34:37.636000',
#  'CnccOrigSendTime': '103436',
#  'MesgRefId': '2018052481056594',
#  'transaction_id': '2018052481056594',
#  'CnccOrigReceiverSID': 'IBPS   ',
#  'Sender': '\xe4\xb8\xad\xe5\x9b\xbd\xe6\xb0\x91\xe7\x94\x9f\xe9\x93\xb6\xe8\xa1\x8c\xe6\x80\xbb\xe8\xa1\x8c',
#  'CnccMesgType': 'ibps.101.001.02   ',
#  'DestIp': '11.194.32.33',
#  'sendTime': '20180524103436',
#  'Receiver': '\xe5\xb9\xb3\xe5\xae\x89\xe9\x93\xb6\xe8\xa1\x8c',
#  'CnccReserve': 'SM2151  97E5C65E                                24T10:34:37.563835  24T10:34:37.647591  MBFEJ               MSGMBFE_1           1527129276          24T10:34:37.636000     24T10:34:37.659326     1005A               MSGCCPC_1           1527129277                           X',
#  'PartId': 1,
#  'sys_type': 'IBPS        ',
#  'TcpSeq': 3717317998,
#  'MesgId': '2018052481056594',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481056594    ',
#  'DecodeId': -966899370,
#  'FlowId': '847981854740062241:-1132081485667368959:0',
#  'CnccMesgID': '2018052481056594    ',
#  'CnccOrigReceiver': '307584007998',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 2531,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 2620957566,
# 'ccpc_sts': '10:34:37.659326',
#  'CnccMesgDirection': 'U',
#  'reverse_intf': 'intf1',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '305100000013',
#  'CCPCId': '1005A               ',
#  'trans_type': '\xe7\xbd\x91\xe9\x93\xb6\xe8\xb4\xb7\xe8\xae\xb0\xe4\xb8\x9a\xe5\x8a\xa1'}

# ,
#  {'HeaderType': 'H',
#  'Prot': 'cncc 2g',
#  'PktLen': 1061,
#  'ts': 1538982486640354000,
#  'PktId': 368746,
#  'SrcIp': '11.194.32.33',
#  'FlowSide': 0,
#  'mbfe_sts': '',
#  'TcpPldLen': 1007,
#  'ms_key': [-966899370,
#  '2018052481056594',
#  'M'],
#  'SrcPort': 41681,
#  'RRB': 'resp',
#  'RRA': 'resp',
#  'ccpc_cts': '',
#  'CnccOrigSendTime': '103437',
#  'MesgRefId': '2018052481056594',
#  'sys_type': 'IBPS',
#  'CnccOrigReceiverSID': 'IBPS',
#  'Sender': '',
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'DestIp': '11.196.163.1',
#  'sendTime': '20180524103437',
#  'Receiver': '',
#  'PartId': 1,
#  'transaction_id': '2018052481056594',
#  'TcpSeq': 421240110,
#  'MesgId': 'NBCD0697000000015428',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481056594    ',
#  'DecodeId': -966899370,
#  'FlowId': '847275006202520321:-6714579453298606079:0',
#  'CnccMesgID': 'NBCD0697000000015428',
#  'CnccOrigReceiver': '',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 1424,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 1572090755,
#  'ccpc_sts': '',
#  'CnccMesgDirection': 'D',
#  'reverse_intf': 'intf13',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '',
#  'CCPCId': '',
#  'trans_type': 'ccms.990.001.01'}

# ,
#  {'HeaderType': 'H',
#  'Prot': 'cncc 2g',
#  'PktLen': 1065,
#  'ts': 1538982498569622000,
#  'PktId': 391283,
#  'SrcIp': '11.194.32.33',
#  'FlowSide': 0,
#  'mbfe_sts': '',
#  'TcpPldLen': 1011,
#  'ms_key': [-966899370,
#  '2018052481056594',
#  'M'],
#  'SrcPort': 49303,
#  'RRB': 'resp',
#  'RRA': 'resp',
#  'ccpc_cts': '',
#  'CnccOrigSendTime': '103438',
#  'MesgRefId': '2018052481056594',
#  'sys_type': 'IBPS',
#  'CnccOrigReceiverSID': 'IBPS',
#  'Sender': '',
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'DestIp': '11.228.35.1',
#  'sendTime': '20180524103438',
#  'Receiver': '',
#  'PartId': 0,
#  'transaction_id': '2018052481056594',
#  'TcpSeq': 1085914746,
#  'MesgId': 'NB4C3F937A0000102865',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481056594    ',
#  'DecodeId': -966899370,
#  'FlowId': '847275006204584705:-4569177180809986047:0',
#  'CnccMesgID': 'NB4C3F937A0000102865',
#  'CnccOrigReceiver': '',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 1424,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 3087022801,
#  'ccpc_sts': '',
#  'CnccMesgDirection': 'D',
#  'reverse_intf': 'intf13',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '',
#  'CCPCId': '',
#  'trans_type': 'ccms.990.001.01'}

# ]

#  """

# # --------------------
# # --------------------

# """ x=[{'HeaderType': 'U',
#  'Prot': 'cncc 2g',
#  'PktLen': 1207,
#  'ts': 1538982371211058000,
#  'PktId': 16989,
#  'SrcIp': '11.196.163.1',
#  'FlowSide': 0,
#  'mbfe_sts': '10:34:28.355643',
#  'TcpPldLen': 1141,
#  'ms_key': [-966899370,
#  '2018052481055659',
#  'M'],
#  'SrcPort': 61516,
#  'RRB': 'req',
#  'RRA': 'req',
#  'ccpc_cts': '10:34:28.355852',
#  'CnccOrigSendTime': '103428',
#  'MesgRefId': '2018052481055659',
#  'transaction_id': '2018052481055659',
#  'CnccOrigReceiverSID': 'IBPS  ',
#  'Sender': '\xe4\xb8\xad\xe5\x9b\xbd\xe6\xb0\x91\xe7\x94\x9f\xe9\x93\xb6\xe8\xa1\x8c\xe6\x80\xbb\xe8\xa1\x8c',
#  'CnccMesgType': 'ibps.101.001.02  ',
#  'DestIp': '11.194.32.35',
#  'sendTime': '20180524103428',
#  'Receiver': '\xe5\x8d\x8e\xe5\xa4\x8f\xe9\x93\xb6\xe8\xa1\x8c\xe8\x82\xa1\xe4\xbb\xbd\xe6\x9c\x89\xe9\x99\x90\xe5\x85\xac\xe5\x8f\xb8\xe6\x80\xbb\xe8\xa1\x8c',
#  'CnccReserve': 'SM2140    EE1FF753                                24T10:34:28.291465  24T10:34:28.355643  MBFEI               MSGMBFE_1           1527129267          24T10:34:28.355852     24T10:34:28.372608     1005A   MSGCCPC_2           1527129268                                                                                                                          X',
#  'PartId': 1,
#  'sys_type': 'IBPS        ',
#  'TcpSeq': 1441874060,
#  'MesgId': '2018052481055659',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481055659    ',
#  'DecodeId': -966899370,
#  'FlowId': '847981854740062243:-1131518535713947647:0',
#  'CnccMesgID': '2018052481055659    ',
#  'CnccOrigReceiver': '304100040000',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 2531,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 1382415482,
#  'ccpc_sts': '10:34:28.372608',
#  'CnccMesgDirection': 'U',
#  'reverse_intf': 'intf1',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '305100000013',
#  'CCPCId': '1005A               ',
#  'trans_type': '\xe7\xbd\x91\xe9\x93\xb6\xe8\xb4\xb7\xe8\xae\xb0\xe4\xb8\x9a\xe5\x8a\xa1'}

# ,
#  {'HeaderType': 'H',
#  'Prot': 'cncc 2g',
#  'PktLen': 1073,
#  'ts': 1538982371431151000,
#  'PktId': 18011,
#  'SrcIp': '11.194.32.35',
#  'FlowSide': 0,
#  'mbfe_sts': '',
#  'TcpPldLen': 1007,
#  'ms_key': [-966899370,
#  '2018052481055659',
#  'M'],
#  'SrcPort': 54426,
#  'RRB': 'resp',
#  'RRA': 'resp',
#  'ccpc_cts': '',
#  'CnccOrigSendTime': '103428',
#  'MesgRefId': '2018052481055659',
#  'sys_type': 'IBPS',
#  'CnccOrigReceiverSID': 'IBPS',
#  'Sender': '',
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'DestIp': '11.196.163.2',
#  'sendTime': '20180524103428',
#  'Receiver': '',
#  'PartId': 0,
#  'transaction_id': '2018052481055659',
#  'TcpSeq': 4075428821,
#  'MesgId': 'NDD0C6394B0000123821',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481055659  ',
#  'DecodeId': -966899370,
#  'FlowId': '847275014792454914:-3127180875121295359:0',
#  'CnccMesgID': 'NDD0C6394B0000123821',
#  'CnccOrigReceiver': '',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 1424,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 2463707827,
#  'ccpc_sts': '',
#  'CnccMesgDirection':'D',
#  'reverse_intf': 'intf13',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '',
#  'CCPCId': '',
#  'trans_type': 'ccms.990.001.01'}

# ,
#  {'HeaderType': 'H',
#  'Prot': 'cncc 2g',
#  'PktLen': 1065,
#  'ts': 1538982375370313000,
#  'PktId': 41272,
#  'SrcIp': '11.194.32.32',
#  'FlowSide': 0,
#  'mbfe_sts': '',
#  'TcpPldLen': 1011,
#  'ms_key': [-966899370,
#  '2018052481055659',
#  'M'],
#  'SrcPort': 32941,
#  'RRB': 'resp',
#  'RRA': 'resp',
#  'ccpc_cts': '',
#  'CnccOrigSendTime': '103429',
#  'MesgRefId': '2018052481055659',
#  'sys_type': 'IBPS',
#  'CnccOrigReceiverSID': 'IBPS',
#  'Sender': '',
#  'CnccMesgType': 'ccms.990.001.01     ',
#  'DestIp': '11.196.163.2',
#  'sendTime': '20180524103429',
#  'Receiver': '',
#  'PartId': 1,
#  'transaction_id': '2018052481055659',
#  'TcpSeq': 2103840628,
#  'MesgId': 'NAC8DEFCAE0000032228',
#  'MetaType': 1,
#  'CnccMesgRefID': '2018052481055659    ',
#  'DecodeId': -966899370,
#  'FlowId': '847275001907553026:-9174670749749739519:0',
#  'CnccMesgID': 'NAC8DEFCAE0000032228',
#  'CnccOrigReceiver': '',
#  'StreamId': 1,
#  'CnccOrigSendDate': '20180524',
#  'DestPort': 1424,
#  'ret_code': '0',
#  'IpProt': 6,
#  'TcpAck': 2355802696,
#  'ccpc_sts': '',
#  'CnccMesgDirection': 'D',
#  'reverse_intf': 'intf13',
#  'sub_trans_type': 'SHANGHAI',
#  'CnccOrigSender': '',
#  'CCPCId': '',
#  'trans_type': 'ccms.990.001.01'}

# ]
# print len(x) """
# # --------------------
# #
# #
# #

# """ from collections import defaultdict
# cache = defaultdict(list)
# SET = set()
# repeat_ms_key_pair = 0
# for i in body_worker:
#     recodr = eval(i)
#     key = tuple(recodr.get('ms_key'))
#     if key in SET:
#         repeat_ms_key_pair+=1
#     else:
#         SET.add(key)
#     cache[key].append(recodr)
# for c in cache:
#     if len(cache[c]) > 4:
#         print "--------------------"
#         print ""
#         print cache[c]
#         print ""
#         print "--------------------"

# print repeat_ms_key_pair
#  """




# """ x=[
#   {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01    ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847258242944999456:-5232327471411494911:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 267943,
#  'CnccMesgID': 'CG19BF59200000107899',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.226',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 46947,
#  'ts': 1538982459338004000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 429145427,
#  'TcpAck': 349895535}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275    ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257491325722656:-2292321339668692991:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 270624,
#  'CnccMesgID': 'CD8B3A6B530000171304',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.51',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 57392,
#  'ts': 1538982459915409000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 440207511,
#  'TcpAck': 893756626}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257491325722657:-2293165764598824959:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 270660,
#  'CnccMesgID': 'CD625BAEE80000174698',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.51',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype':['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.33',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 57389,
#  'ts': 1538982459919236000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 1703289403,
#  'TcpAck': 2170226436}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847258242944999457:-5232045996434784255:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 268094,
#  'CnccMesgID': 'CG903BB6CC0000110803',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.226',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.33',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 46948,
#  'ts': 1538982459416639000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 498199658,
#  'TcpAck': 2864153988}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257482735788064:-655825825072939007:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 275992,
#  'CnccMesgID': 'CB5494C6D30000325270',
#  'StreamId': 1,
#  'SrcIp':'11.194.16.49',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 63206,
#  'ts': 1538982460725167000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type':'\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 1214515041,
#  'TcpAck': 2425708915}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257482735788065:-655544350096228351:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 276071,
#  'CnccMesgID': 'CB749BC6B50000327636',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.49',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
# 'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.33',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 63207,
#  'ts': 1538982460732400000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 2006000560,
#  'TcpAck': 92951904}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257478440820768:-4365384543142674431:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 286176,
#  'CnccMesgID': 'CAEF0951B50000595491',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.48',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 50027,
#  'ts': 1538982464036821000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 1568445833,
#  'TcpAck': 2851789556}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 185,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847257478440820769:-4363977168259121151:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 286268,
#  'CnccMesgID': 'CA4F53EA710000598216',
#  'StreamId': 1,
#  'SrcIp': '11.194.16.48',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 131,
#  'PartId': 0,
#  'DestIp': '11.194.32.33',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 50032,
#  'ts': 1538982464041332000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
# 'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 725400381,
#  'TcpAck': 2670360943}
# ,
#  {'CnccOrigReceiverSID': 'PMTS',
#  'status': None,
#  'ms_key': [-971229470,
#  '1527129275          ',
#  'M'],
#  'MetaType': 1,
#  'CnccMesgType': 'ccms.992.001.01     ',
#  'CnccMesgRefID': '1527129275          ',
#  'Prot': 'xml',
#  'PktLen': 78,
#  'DecodeId': -971229470,
#  'reverse_intf': 'intf10',
#  'RRB': 'ccms.992.001.01',
#  'FlowId': '847274971842617376:-8340937114203979775:0',
#  'sub_trans_type': 'ccms',
#  'PktId': 301794,
#  'CnccMesgID': 'CFA5DB48620000999034',
#  'StreamId': 1,
#  'SrcIp': '11.194.32.25',
#  'FlowSide': 0,
#  'ibpstype': ['ibps.311',
#  'ibps.309',
#  'ibps.305',
#  'ibps.331',
#  'ibps.333',
#  'ibps.335',
#  'ibps.337',
#  'ibps.339',
#  'ibps.341',
#  'ibps.343',
#  'ibps.703',
#  'ibps.705'],
#  'DestPort': 2531,
#  'ccmstype': ['ccms.303',
#  'ccms.307',
#  'ccms.308',
#  'ccms.310',
#  'ccms.311',
#  'ccms.312',
#  'ccms.313',
#  'ccms.314',
#  'ccms.315',
#  'ccms.316',
#  'ccms.317',
#  'ccms.318',
#  'ccms.319',
#  'ccms.801',
#  'ccms.803',
#  'ccms.805',
#  'ccms.806',
#  'ccms.807',
#  'ccms.809',
#  'ccms.811',
#  'ccms.900',
#  'ccms.903',
#  'ccms.906',
#  'ccms.907',
#  'ccms.911',
#  'ccms.913',
#  'ccms.915',
#  'ccms.916',
#  'ccms.917',
#  'ccms.919',
#  'ccms.921',
#  'ccms.926',
#  'ccms.990',
#  'ccms.991',
#  'ccms.992'],
#  'sys_type': 'PMTS',
#  'TcpPldLen': 24,
#  'PartId': 0,
#  'DestIp': '11.194.32.32',
#  'mesgtype': 'ccms.992',
#  'ret_code': None,
#  'SrcPort': 35903,
#  'ts': 1538982466799091000,
#  'RRA': 'ccms.992.001.01',
#  'transaction_id': '1527129275          ',
#  'IpProt': 6,
#  'trans_type': '\xe6\x8e\xa2\xe6\xb5\x8b\xe5\x9b\x9e\xe5\xba\x94',
#  'TcpSeq': 3438100462,
#  'TcpAck': 3240370874}
# ]



# print len(x) """


# """ sets = set()

# for i in body_worker:
#     recodr = eval(i)
#     rrb = recodr.get('RRB')
#     if rrb != 'req' and rrb != 'resp' and rrb !='notify':
#         if rrb not in sets:
#             sets.add(rrb)
#             print recodr
# for i in sets:
#     print i

#  """
# """ illegal_rrb = 0

# for i in body_worker:
#     recodr = eval(i)
#     rrb = recodr.get('RRB')
#     if rrb != 'req' and rrb != 'resp' and rrb !='notify':
#         illegal_rrb+=1
# print illegal_rrb

#  """


# # 第一次
# #  251 界面发包的交易量为 16176
# #  251 worker   body 里面的行数 为 49150 
# #  251 switch   body 里面的行数 为 49158
# #  251 exporter body 里面的行数为 20121





# # CnccOrigReceiverSID  HVPS 

# #
# with open('/home/chauncey/Desktop/tmp/bodyfile_2018_10_10',
# 'r') as f:
#     body_worker = f.readlines()


# HVPS_IP_LIST = [
#     "11.194.16.16",

#     "11.194.16.17",

#     "11.194.16.18",

#     "11.194.16.19",

# ]

# # CnccOrigReceiverSID =  "HVPS"
# # CnccOrigReceiverSID =  "HVPS"

# # sys_type = "HVPS"

# RRA = "req"
# NOTIFY = "notify"
# hvpd_sum = 0 

# # for i in body_worker:
# #     recodr = eval(i)
# #     if recodr.get("DestIp") in HVPS_IP_LIST  and recodr.get("RRA") == [RRA,
# NOTIFY]:
# #         hvpd_sum +=1 
# #         if hvpd_sum == 4:
# #             print recodr
# # print hvpd_sum
# #1586 
# # 1766 sys_type DestIp  RRA
# # 1774 DestIp RRA



# # NPC-->HVPS
# # hvps.111.001.01--1571  req
# # hvps.112.001.01--1    req
# # hvps.141.001.01--6      req 
# # ccms.314.001.01--4      req 
# # ccms.315.001.01--4      resp 
# # ccms.991.001.01--28      req
# # ccms.992.001.01--16      resp
# # hvps 总交易量 


# # S = set()

# # for i in body_worker:
# #     recodr = eval(i)
# #     if recodr.get('RRA') == NOTIFY:
# #         S.add(recodr.get("ms_key"))

# # for i in body_worker:
# #     recodr = eval(i)
# #     if recodr.get('RRA') != NOTIFY:
# #         if recodr.get("ms_key") in S:
# #             print recodr


# reqSUM = 0
# for i in body_worker:
#     recodr = eval(i)
#     if recodr.get("RRA")== "req" :
#         reqSUM +=1
# print reqSUM

# trans_count = 0 
# for i in body_worker:
#     record = eval(i)
#     trans_count += int( record.get('trans_count',0))
# print trans_count



""" {'ts': 1539155191230240000,
 'FlowId': '847275010497318928:-1827606153119399935:0',
 'PktLen': 126,
 'MetaType': 1,
 'StreamId': 1,
 'PartId': 1,
 'DecodeId': 1825566008,
 'FlowSide': 0,
 'SrcIp': '11.194.32.34',
 'DestIp': '11.194.16.16',
 'SrcPort': 59043,
 'DestPort': 2531,
 'IpProt': 6,
 'TcpPldLen': 72,
 'PktId': 15817125,
 'TcpSeq': 2054667427,
 'TcpAck': 1077684628,
 'CnccOrigReceiverSID': 'HVPS',
 'CnccMesgType': 'ccms.303.001.02     ',
 'CnccMesgID': '20180611000275257366',
 'CnccMesgRefID': '20180611000098019594',
 'MsgId': '2018061120372042',
 'RRA': 'notify',
 'Prot': 'xml',
 'RRB': 'notify',
 'mesgtype': 'ccms.303',
 'sub_trans_type': 'ccms',
 'transaction_id': '2018061120372042',
 'transactionId': '2018061120372042',
 'ccmstype': ('ccms.303',
 'ccms.307',
 'ccms.308',
 'ccms.310',
 'ccms.311',
 'ccms.312',
 'ccms.313',
 'ccms.314',
 'ccms.315',
 'ccms.316',
 'ccms.317',
 'ccms.318',
 'ccms.319',
 'ccms.801',
 'ccms.803',
 'ccms.805',
 'ccms.806',
 'ccms.807',
 'ccms.809',
 'ccms.811',
 'ccms.900',
 'ccms.903',
 'ccms.906',
 'ccms.907',
 'ccms.911',
 'ccms.913',
 'ccms.915',
 'ccms.916',
 'ccms.917',
 'ccms.919',
 'ccms.921',
 'ccms.926',
 'ccms.990',
 'ccms.991',
 'ccms.992'),
 'trans_type': '\xe8\x87\xaa\xe7\x94\xb1\xe6\xa0\xbc\xe5\xbc\x8f\xe6\x8a\xa5\xe6\x96\x87',
 'ret_code': None,
 'status': '--',
 'sys_type': 'HVPS',
'2018061120372042',
 'M') 'ms_key': (1825566008,
 '2018061120372042',
 'M'),
 'reverse_intf': u'intf3'}





 """


# DecodeId
# import os,sys,msgpack,time
# def unpack_raw(stream):
#     unpacker = msgpack.Unpacker(stream)
#     for msg in unpacker:
#         yield msg
# if len(sys.argv) != 2:
#     print " unpack.py dirname"

# # 20181017103700_20181017103700_0.btr.pack
# #ts_start = 20181017103700
# #ts_end = 20181017103700 + 60
# def format_time(t):
#     return time.mktime(time.strptime(t,"%Y%m%d%H%M%M"))
# packDir = sys.argv[1]
# if packDir.endswith('/'):
#     packDir=packDir.rstrip("/")
# for packFileName in os.listdir(packDir):
#     ts_start = format_time(int(packFileName.split('_')[0]))
#     ts_end = format_time(ts_start+60)
#     for msg in unpack_raw(file(packDir+os.sep+packFileName)):
#         print packDir+os.sep+packFileName
#         record = eval(msg)
#         ts = int(record.get('ts'))
#         ts = int(ts / 10000000)
#         if ts >=ts_start and ts <= ts_end:
#             pass
#     else:
#         print record
# for i in body_worker:
#     record = eval(i)
#     ts = int(record.get('ts'))
#     ts = int(ts / 1000000000)
#     if ts >=1539743820 and ts <= 1539743880:
#         pass
#     else:
#         print record






['/opt/cbms/var/store/btr/app2/intf5/20181017/20181017173100_20181017173100_1.btr.pack',
 '/opt/cbms/var/store/btr/app2/intf5/20181017/20181017173100_20181017173100_0.btr.pack']









