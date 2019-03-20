import xmltodict
import csv
# csv.reader()
# with open("cnccU_a.xml") as f:
#     xml_format= "<dpxml>{}</dpxml>".format(f.read())
#     x=xmltodict.parse(xml_format)
#     print(x)

x = {'@id': '1', '@baseProtocol': '', 'property': {'attribute': [{'@key': 'forceCheckPayload', '@value': 'a=wmq,b=http,c=dghk', '@type': 'Boolean'}]}, 'recordField': {'@allItem': 'false', '@template': 'xpath', 'field': [{'@item': '/Document/*/TxInfAndSts/StsId', '#text': 'StsId'}, {'@item': '/Document/*/NPCPrcInf/PrcSts', '#text': 'PrcSts'}, {'@item': '/Document/*/RspnInf/PrcCd', '#text': 'PrcCd'}, {'@item': '/Document/*/DscrdInf/PrcCd', '#text': 'PrcCd'}, {'@item': '/Document/*/CdtTrfTxInf/PmtTpInf/*/Prtry', '#text': 'prtry1'}, {'@item': '/Document/*/GrpHdr/MsgId', '#text': 'MsgId'}, {'@item': '/Document/*/DscrdInf/MsgId', '#text': 'MsgId'}, {'@item': '/Document/*/CdtTrfTxInf/PmtId/TxId', '#text': 'TxId'}, {'@item': '/Document/*/OrgnlGrpHdr/OrgnlMsgId', '#text': 'OrgnlMsgId'}, {'@item': '/Document/*/OrgnlGrpInfAndSts/OrgnlMsgId', '#text': 'OrgnlMsgId'}, {'@item': '/Document/*/DscrdInf/MsgRefId', '#text': 'MsgRefId'}, {'@item': '/Document/*/WdrwlRspnInf/OrgnlTxId', '#text': 'OrgnlTxId'}, {'@item': '/Document/SysStsNtfctn/SysStsInf/OrgnlSysSts', '#text': 'org_status'}, {'@item': '/Document/SysStsNtfctn/SysStsInf/CurSysSts', '#text': 'sys_status'}]}}
print(xmltodict.unparse({"protocol":x}))
# p = {
#     "protocol": {
#         "dp.xml": [
#             {
#                 "@id": "elit culpa magna Duis",
#                 "@baseProtocol": "et mollit",
#                 "recordField": {
#                     "@allItem": "occaecat labore",
#                     "@template": "tempor Lorem minim aute",
#                     "field": [
#                         {
#                             "@item": "adipisicing laboris sit Ut dolor",
#                             "#text": "commodo fugiat magna dolor quis"
#                         },
#                         {
#                             "@item": "irure ad velit",
#                             "#text": "ea ut mollit consequat"
#                         },
#                         {
#                             "@item": "ex",
#                             "#text": "minim exercitation mollit tempor"
#                         }
#                     ]
#                 }
#             },
#             {
#                 "@id": "d",
#                 "@baseProtocol": "tempor",
#                 "recordField": {
#                     "@allItem": "dolore",
#                     "@template": "non in sed exercitation magna",
#                     "field": [
#                         {
#                             "@item": "est sit mollit occaecat",
#                             "#text": "ipsum commodo culpa nulla id"
#                         },
#                         {
#                             "@item": "adipisicing su",
#                             "#text": "quis"
#                         }
#                     ]
#                 }
#             },
#             {
#                 "@id": "eu mollit dolore ullamco ",
#                 "@baseProtocol": "minim sit",
#                 "recordField": {
#                     "@allItem": "qui Excepteur reprehenderit",
#                     "@template": "consectetur ut Duis e",
#                     "field": [
#                         {
#                             "@item": "et ut",
#                             "#text": "dolor magna sunt sint"
#                         },
#                         {
#                             "@item": "mollit dolore",
#                             "#text": "commodo quis"
#                         },
#                         {
#                             "@item": "eu voluptate",
#                             "#text": "anim consectetur commodo Ut"
#                         },
#                         {
#                             "@item": "qui dolor",
#                             "#text": "est dolor et"
#                         }
#                     ]
#                 }
#             },
#             {
#                 "@id": "adipisicing ipsum do exercitation",
#                 "@baseProtocol": "sint Lorem nisi ad ",
#                 "recordField": {
#                     "@allItem": "ex est",
#                     "@template": "do",
#                     "field": [
#                         {
#                             "@item": "deserunt cillum culpa aute",
#                             "#text": "esse"
#                         },
#                         {
#                             "@item": "nulla",
#                             "#text": "pariatur dolore eu"
#                         },
#                         {
#                             "@item": "ipsum",
#                             "#text": "dolor ipsum sint"
#                         },
#                         {
#                             "@item": "laborum proident aliquip",
#                             "#text": "nisi magna ut non"
#                         }
#                     ]
#                 }
#             },
#             {
#                 "@id": "ipsum in sint",
#                 "@baseProtocol": "sunt officia aliquip voluptate ex",
#                 "recordField": {
#                     "@allItem": "dolore velit consequat sint Excepteur",
#                     "@template": "laboris ut elit e",
#                     "field": [
#                         {
#                             "@item": "reprehenderit",
#                             "#text": "irure"
#                         },
#                         {
#                             "@item": "pariatur in dolor aliquip",
#                             "#text": "sed reprehenderit"
#                         }
#                     ]
#                 }
#             }
#         ],
#         "proto.xml": {
#             "basic": {
#                 "name": "sit",
#                 "display_name": "qui et irure",
#                 "decodes": [
#                     "officia id tempo"
#                 ]
#             },
#             "prepare": [
#                 {
#                     "@op": "ea amet nostrud proident culpa",
#                     "@value": "aute tempor i",
#                     "#text": "culpa reprehenderit",
#                     "@item1": "eiusmod ad",
#                     "@start": "amet sint",
#                     "@end": "ex"
#                 },
#                 {
#                     "@op": "culpa incididunt sed ex",
#                     "@value": "cupidatat",
#                     "#text": "ad ut sint",
#                     "@item1": "sint magna consectetur ad",
#                     "@start": "et Ut magna",
#                     "@end": "ut"
#                 },
#                 {
#                     "@op": "ad incididunt dolor",
#                     "@value": "aute laboris consequat dolore",
#                     "#text": "est aliquip",
#                     "@item1": "ea aliqua",
#                     "@start": "id in laboris minim",
#                     "@end": "id Lorem aute"
#                 }
#             ],
#             "normalizes": [
#                 {
#                     "@name": "labore officia",
#                     "if": [
#                         {
#                             "@op": "ex in",
#                             "@item1": "ad cillum laborum do",
#                             "field": [
#                                 {
#                                     "@op": "et ",
#                                     "@item1": "enim labore pariatur proident",
#                                     "@item2": "consequat ullamco Excepteur",
#                                     "@start": "co",
#                                     "@end": "dolor esse cupidatat"
#                                 },
#                                 {
#                                     "@op": "adipisicing sint laboris sed",
#                                     "@item1": "est aute mollit",
#                                     "@item2": "ut in",
#                                     "@start": "irure est proid",
#                                     "@end": "culpa minim eiusmod ad occaecat"
#                                 },
#                                 {
#                                     "@op": "dolor reprehenderit",
#                                     "@item1": "culpa voluptate incididunt",
#                                     "@item2": "dolore adipisicing veniam",
#                                     "@start": "laborum id proident elit",
#                                     "@end": "velit aliqua magna Ut"
#                                 },
#                                 {
#                                     "@op": "irure",
#                                     "@item1": "nisi consectetur nulla",
#                                     "@item2": "commodo dolor esse Duis",
#                                     "@start": "enim ut ex amet sunt",
#                                     "@end": "officia elit"
#                                 }
#                             ],
#                             "@item2": "veniam incididunt ut Ut"
#                         },
#                         {
#                             "@op": "Lorem est",
#                             "@item1": "ipsum fugiat consequat",
#                             "field": [
#                                 {
#                                     "@op": "cupidatat dolor eiusmod amet",
#                                     "@item1": "laborum do",
#                                     "@item2": "aute cillum ipsum reprehenderit",
#                                     "@start": "non dolo",
#                                     "@end": "enim"
#                                 },
#                                 {
#                                     "@op": "in dolore id pariatur dolor",
#                                     "@item1": "aliquip elit ad",
#                                     "@item2": "laboris anim in cillum",
#                                     "@start": "ex deserunt occaecat",
#                                     "@end": "aliquip est la"
#                                 },
#                                 {
#                                     "@op": "laboris ut cupidatat esse",
#                                     "@item1": "aliqua reprehenderit adipisicing",
#                                     "@item2": "Ut velit qui in ",
#                                     "@start": "incididunt velit non",
#                                     "@end": "sint nisi et adipisicing quis"
#                                 },
#                                 {
#                                     "@op": "magna ut",
#                                     "@item1": "nostrud ut ad minim",
#                                     "@item2": "non",
#                                     "@start": "laborum aliqua fugiat mollit",
#                                     "@end": "ad labor"
#                                 }
#                             ],
#                             "@item2": "dolore reprehenderit"
#                         },
#                         {
#                             "@op": "in sed",
#                             "@item1": "in anim amet veniam",
#                             "field": [
#                                 {
#                                     "@op": "in cu",
#                                     "@item1": "tempor magna id",
#                                     "@item2": "elit occaecat consectetur non",
#                                     "@start": "consequat ea Lorem sit",
#                                     "@end": "elit dolore anim"
#                                 },
#                                 {
#                                     "@op": "sunt elit fugiat consectetur",
#                                     "@item1": "dolor magna anim tempor amet",
#                                     "@item2": "esse dolor Excepteur et anim",
#                                     "@start": "occaecat anim consectetur",
#                                     "@end": "dolor Lorem nostrud mollit"
#                                 },
#                                 {
#                                     "@op": "tempor",
#                                     "@item1": "Duis eiusmod nisi veniam nostrud",
#                                     "@item2": "pariatur amet sunt labore",
#                                     "@start": "dolor",
#                                     "@end": "ad in dolore"
#                                 },
#                                 {
#                                     "@op": "Duis in laboris quis",
#                                     "@item1": "sint laborum",
#                                     "@item2": "aute enim laboris",
#                                     "@start": "sunt ipsum pariatur",
#                                     "@end": "Duis"
#                                 },
#                                 {
#                                     "@op": "sunt dolore exercitation minim",
#                                     "@item1": "nostrud ipsu",
#                                     "@item2": "exercitation adipisicing",
#                                     "@start": "cillum exercitation reprehenderit",
#                                     "@end": "deserunt proident"
#                                 }
#                             ],
#                             "@item2": "aliqua Excepteur "
#                         },
#                         {
#                             "@op": "Ut adipisicing Lorem dolor",
#                             "@item1": "nisi sit elit ullamco ut",
#                             "field": [
#                                 {
#                                     "@op": "dolor mollit com",
#                                     "@item1": "Lorem ea nostrud",
#                                     "@item2": "amet commodo esse id",
#                                     "@start": "consequat dolore",
#                                     "@end": "pariatur dolor do Excepteur eiusmod"
#                                 },
#                                 {
#                                     "@op": "sunt",
#                                     "@item1": "magna Lorem reprehenderit",
#                                     "@item2": "qui in elit",
#                                     "@start": "fugiat amet ex officia ipsum",
#                                     "@end": "ea mollit"
#                                 },
#                                 {
#                                     "@op": "labore officia",
#                                     "@item1": "sunt",
#                                     "@item2": "dolor irure proident id",
#                                     "@start": "incididunt",
#                                     "@end": "ipsum amet adipisicing"
#                                 },
#                                 {
#                                     "@op": "minim eu pariatur cupidatat",
#                                     "@item1": "sed laboris ut",
#                                     "@item2": "ex",
#                                     "@start": "ipsum consequat Excepteur c",
#                                     "@end": "dolore sed amet"
#                                 }
#                             ],
#                             "@item2": "mollit nostrud"
#                         },
#                         {
#                             "@op": "ullamco ut eiusmod",
#                             "@item1": "ut laboris",
#                             "field": [
#                                 {
#                                     "@op": "dolor et magna ",
#                                     "@item1": "veniam velit",
#                                     "@item2": "est sunt",
#                                     "@start": "ullamco deserunt",
#                                     "@end": "ut officia sed exercitation"
#                                 }
#                             ],
#                             "@item2": "Duis vo"
#                         }
#                     ],
#                     "else": [
#                         {
#                             "@op": "irure in nostrud",
#                             "@value": "laboris cupidatat",
#                             "@item1": "in occaecat",
#                             "@item2": "non laborum ea",
#                             "@start": "amet elit sunt",
#                             "@end": "ullamco non aliquip"
#                         },
#                         {
#                             "@op": "dolore eu dolor in",
#                             "@value": "Ut quis in",
#                             "@item1": "eiusmod elit anim ea non",
#                             "@item2": "e",
#                             "@start": "Excepteur amet",
#                             "@end": "sit elit qui"
#                         }
#                     ]
#                 },
#                 {
#                     "@name": "velit incididu",
#                     "if": [
#                         {
#                             "@op": "adipisicing do Excepteur in",
#                             "@item1": "sit do id adipisicing",
#                             "field": [
#                                 {
#                                     "@op": "enim fugiat velit",
#                                     "@item1": "amet magna id Ut",
#                                     "@item2": "ut Lorem aliquip",
#                                     "@start": "aliquip dolor fugiat",
#                                     "@end": "adipisicing id est Excepteur"
#                                 },
#                                 {
#                                     "@op": "Duis ut",
#                                     "@item1": "mollit dolore Ut aliqua dolor",
#                                     "@item2": "adipisicing eu velit nulla",
#                                     "@start": "in ex exercitation",
#                                     "@end": "non dolore in"
#                                 }
#                             ],
#                             "@item2": "vel"
#                         },
#                         {
#                             "@op": "dolore dolor",
#                             "@item1": "aliquip sunt",
#                             "field": [
#                                 {
#                                     "@op": "Excepteur mollit dolore",
#                                     "@item1": "Duis",
#                                     "@item2": "sit et consectetur non",
#                                     "@start": "cupidatat in irure eu",
#                                     "@end": "Lorem"
#                                 },
#                                 {
#                                     "@op": "in quis Duis sit",
#                                     "@item1": "ipsum",
#                                     "@item2": "deserunt labori",
#                                     "@start": "ut et adipisicing",
#                                     "@end": "nulla in velit elit"
#                                 },
#                                 {
#                                     "@op": "occaecat aliquip incididunt",
#                                     "@item1": "ex ipsum irure enim",
#                                     "@item2": "proident esse",
#                                     "@start": "Ut irure",
#                                     "@end": "officia id anim cillum ex"
#                                 }
#                             ],
#                             "@item2": "ut labore commodo"
#                         }
#                     ],
#                     "else": [
#                         {
#                             "@op": "exercitation dolor est",
#                             "@value": "veniam eu proident in",
#                             "@item1": "ut irure magna occaecat",
#                             "@item2": "id do Excepteur tempor commodo",
#                             "@start": "enim eu ullamco",
#                             "@end": "enim sint dolor"
#                         },
#                         {
#                             "@op": "do Lorem",
#                             "@value": "nisi sint do",
#                             "@item1": "exercitation in adipisicing",
#                             "@item2": "qui ipsum",
#                             "@start": "irure tempor qui",
#                             "@end": "reprehenderit quis non"
#                         },
#                         {
#                             "@op": "ea commodo",
#                             "@value": "fugiat commodo est anim do",
#                             "@item1": "elit occaecat consectetur fugiat deserunt",
#                             "@item2": "re",
#                             "@start": "ipsum amet in anim",
#                             "@end": "veniam Duis Ut velit ea"
#                         },
#                         {
#                             "@op": "sed eu pariatur",
#                             "@value": "occaecat",
#                             "@item1": "dolore fugiat cillum",
#                             "@item2": "eu ullamco ad",
#                             "@start": "laboris cupidatat consequat",
#                             "@end": "sit amet aute pariatur"
#                         }
#                     ]
#                 }
#             ],
#             "traces": [
#                 {
#                     "@name": "amet enim sed mollit",
#                     "field": [
#                         {
#                             "@op": "aliquip",
#                             "@item1": "velit laboris ut"
#                         },
#                         {
#                             "@op": "non ad velit laboris",
#                             "@item1": "occaecat eu"
#                         },
#                         {
#                             "@op": "dolore",
#                             "@item1": "laboris exercitation voluptate"
#                         },
#                         {
#                             "@op": "nulla dolor anim con",
#                             "@item1": "labore"
#                         },
#                         {
#                             "@op": "est cillum ea voluptate",
#                             "@item1": "in pariatur in in fugiat"
#                         }
#                     ]
#                 },
#                 {
#                     "@name": "pariatur ullamco aute proident",
#                     "field": [
#                         {
#                             "@op": "aute irure",
#                             "@item1": "enim"
#                         },
#                         {
#                             "@op": "anim ut nostrud nulla",
#                             "@item1": "laborum"
#                         },
#                         {
#                             "@op": "dolore",
#                             "@item1": "non cupidatat ad ut"
#                         }
#                     ]
#                 },
#                 {
#                     "@name": "fugiat",
#                     "field": [
#                         {
#                             "@op": "cillum ut exercitation non tempor",
#                             "@item1": "incididunt sunt laboris veniam cupidatat"
#                         },
#                         {
#                             "@op": "qui deserunt cillum",
#                             "@item1": "enim in reprehenderit in"
#                         },
#                         {
#                             "@op": "elit Duis sed nostrud",
#                             "@item1": "sunt dolore laboris in"
#                         },
#                         {
#                             "@op": "irure in magna nulla",
#                             "@item1": "magna culpa minim"
#                         },
#                         {
#                             "@op": "adipisicing",
#                             "@item1": "ea quis sint aliqua"
#                         }
#                     ]
#                 }
#             ],
#             "trace_disp": [
#                 {
#                     "@op": "reprehenderit do est amet exercitation",
#                     "@item1": "",
#                     "#text": "sint adipisicing consectetur"
#                 },
#                 {
#                     "@op": "id eiusmod minim non officia",
#                     "@item1": "qui in cupid",
#                     "#text": "laborum"
#                 }
#             ],
#             "dimensions": [
#                 {
#                     "@name": "Ut minim ",
#                     "@display_name": "labore pariatur",
#                     "@csv_file": "qui Lorem"
#                 },
#                 {
#                     "@name": "do cupidatat",
#                     "@display_name": "non ipsum ut eu",
#                     "@csv_file": "irure"
#                 }
#             ]
#         }
#     }
# }
# x = {
#     "protocol": {
#         "dp.xml": [
#             {
#                 "@id": "xml",
#                 "@baseProtocol": "cncc 2g",
#                 "recordField": {
#                     "@allItem": "false",
#                     "@template": "xpath",
#                     "field": [
#                         {
#                             "@item": "/Document/*/TxInfAndSts/StsId",
#                             "#text": "StsId"
#                         },
#                         {
#                             "@item": "/Document/*/NPCPrcInf/PrcSts",
#                             "#text": "PrcSts"
#                         },
#                         {
#                             "@item": "/Document/*/RspnInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtTpInf/*/Prtry",
#                             "#text": "prtry1"
#                         },
#                         {
#                             "@item": "/Document/*/GrpHdr/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtId/TxId",
#                             "#text": "TxId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpHdr/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpInfAndSts/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgRefId",
#                             "#text": "MsgRefId"
#                         },
#                         {
#                             "@item": "/Document/*/WdrwlRspnInf/OrgnlTxId",
#                             "#text": "OrgnlTxId"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/OrgnlSysSts",
#                             "#text": "org_status"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/CurSysSts",
#                             "#text": "sys_status"
#                         }
#                     ]
#                 }
#             }
#         ],
#         "proto.xml": {
#             "basic": {
#                 "name": "cnccU",
#                 "display_name": "人行清算中心U头报文",
#                 "decodes": {
#                     "decode": [
#                         "1827731058",
#                         "1825566008",
#                         "1832061158"
#                     ]
#                 }
#             },
#             "prepare": {
#                 "filed": [
#                     {
#                         "@op": "as",
#                         "@value": "0",
#                         "#text": "ret_code"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "1",
#                         "#text": "is_success"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "120",
#                         "#text": "max_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "20",
#                         "#text": "overtime_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "60",
#                         "#text": "statistics_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "60",
#                         "#text": "delay_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "1",
#                         "#text": "alert_threshold"
#                     },
#                     {
#                         "@op": "as",
#                         "@item1": "CnccOrigReceiverSID",
#                         "#text": "sys_type"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgType",
#                         "#text": "CnccMesgType"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgType",
#                         "#text": "RRB"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgID",
#                         "#text": "MesgId"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgRefID",
#                         "#text": "MesgRefId"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "196",
#                         "@end": "216",
#                         "#text": "CCPCId"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "73",
#                         "@end": "88",
#                         "#text": "mbfe_sts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "153",
#                         "@end": "168",
#                         "#text": "ccpc_cts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "176",
#                         "@end": "191",
#                         "#text": "ccpc_sts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "DestIp",
#                         "@start": "3",
#                         "@end": "6",
#                         "#text": "sub_trans_type"
#                     }
#                 ]
#             },
#             "normalizes": {
#                 "normalize": [
#                     {
#                         "@name": "CnccOrigSender",
#                         "if": [
#                             {
#                                 "@op": "notNull",
#                                 "@item1": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "notNull",
#                                 "@value": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "eq",
#                                 "@item1": "CnccOrigSender",
#                                 "@item2": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             }
#                         ],
#                         "else": {
#                             "filed": [
#                                 {
#                                     "@op": "as",
#                                     "@value": "--"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_add",
#                                     "@item1": "CnccOrigReceiver",
#                                     "@item2": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_sub",
#                                     "@item1": "CnccReserve",
#                                     "@start": "176",
#                                     "@end": "191"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 }
#                             ]
#                         }
#                     },
#                     {
#                         "@name": "CnccOrigSender",
#                         "if": [
#                             {
#                                 "@op": "notNull",
#                                 "@item1": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "eq",
#                                 "@item1": "CnccOrigSender",
#                                 "@item2": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             }
#                         ],
#                         "else": {
#                             "filed": [
#                                 {
#                                     "@op": "as",
#                                     "@value": "--"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_add",
#                                     "@item1": "CnccOrigReceiver",
#                                     "@item2": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_sub",
#                                     "@item1": "CnccReserve",
#                                     "@start": "176",
#                                     "@end": "191"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 }
#                             ]
#                         }
#                     },
#                     {
#                         "filed": [
#                             {
#                                 "@op": "as",
#                                 "@value": "--"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_add",
#                                 "@item1": "CnccOrigReceiver",
#                                 "@item2": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_sub",
#                                 "@item1": "CnccReserve",
#                                 "@start": "176",
#                                 "@end": "191"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             }
#                         ]
#                     }
#                 ]
#             },
#             "traces": {
#                 "filed": [
#                     {
#                         "@name": "critical_keyword",
#                         "field": {
#                             "@op": "as",
#                             "@item1": "transaction_id"
#                         }
#                     },
#                     {
#                         "@name": "fields",
#                         "field": [
#                             {
#                                 "@op": "as",
#                                 "@item1": "status"
#                             },
#                             {
#                                 "@op": "as",
#                                 "@item1": "TxId"
#                             }
#                         ]
#                     }
#                 ]
#             },
#             "trace_disp": {
#                 "filed": [
#                     {
#                         "@op": "disp",
#                         "@item1": "status",
#                         "#text": "业务状态"
#                     },
#                     {
#                         "@op": "disp",
#                         "@item1": "TxId",
#                         "#text": "明细标识号"
#                     }
#                 ]
#             },
#             "dimensions": {
#                 "dimension": [
#                     {
#                         "@name": "trans_type",
#                         "@display_name": "业务类型",
#                         "@csv_file": "trans_type.csv"
#                     },
#                     {
#                         "@name": "sub_trans_type",
#                         "@display_name": "业务类型",
#                         "@csv_file": "sub_trans_type.csv"
#                     }
#                 ]
#             }
#         }
#     }
# }
# p = {
#     "protocol": {
#         "dp.xml": [
#             {
#                 "@id": "tcp",
#                 "@baseProtocol": "",
#                 "property": {
#                     "attribute": [
#                         {
#                             "@key": "forceCheckPayload",
#                             "@value": "true",
#                             "@type": "Boolean"
#                         }
#                     ]
#                 },
#                 "recordField": {
#                     "@allItem": "false",
#                     "@template": "xpath",
#                     "field": [
#                         {
#                             "@item": "/Document/*/TxInfAndSts/StsId",
#                             "#text": "StsId"
#                         },
#                         {
#                             "@item": "/Document/*/NPCPrcInf/PrcSts",
#                             "#text": "PrcSts"
#                         },
#                         {
#                             "@item": "/Document/*/RspnInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtTpInf/*/Prtry",
#                             "#text": "prtry1"
#                         },
#                         {
#                             "@item": "/Document/*/GrpHdr/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtId/TxId",
#                             "#text": "TxId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpHdr/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpInfAndSts/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgRefId",
#                             "#text": "MsgRefId"
#                         },
#                         {
#                             "@item": "/Document/*/WdrwlRspnInf/OrgnlTxId",
#                             "#text": "OrgnlTxId"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/OrgnlSysSts",
#                             "#text": "org_status"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/CurSysSts",
#                             "#text": "sys_status"
#                         }
#                     ]
#                 }
#             },
#             {
#                 "@id": "xml",
#                 "@baseProtocol": "cncc 2g",
#                 "recordField": {
#                     "@allItem": "false",
#                     "@template": "xpath",
#                     "field": [
#                         {
#                             "@item": "/Document/*/TxInfAndSts/StsId",
#                             "#text": "StsId"
#                         },
#                         {
#                             "@item": "/Document/*/NPCPrcInf/PrcSts",
#                             "#text": "PrcSts"
#                         },
#                         {
#                             "@item": "/Document/*/RspnInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/PrcCd",
#                             "#text": "PrcCd"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtTpInf/*/Prtry",
#                             "#text": "prtry1"
#                         },
#                         {
#                             "@item": "/Document/*/GrpHdr/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgId",
#                             "#text": "MsgId"
#                         },
#                         {
#                             "@item": "/Document/*/CdtTrfTxInf/PmtId/TxId",
#                             "#text": "TxId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpHdr/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/OrgnlGrpInfAndSts/OrgnlMsgId",
#                             "#text": "OrgnlMsgId"
#                         },
#                         {
#                             "@item": "/Document/*/DscrdInf/MsgRefId",
#                             "#text": "MsgRefId"
#                         },
#                         {
#                             "@item": "/Document/*/WdrwlRspnInf/OrgnlTxId",
#                             "#text": "OrgnlTxId"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/OrgnlSysSts",
#                             "#text": "org_status"
#                         },
#                         {
#                             "@item": "/Document/SysStsNtfctn/SysStsInf/CurSysSts",
#                             "#text": "sys_status"
#                         }
#                     ]
#                 }
#             }
#         ],
#         "proto.xml": {
#             "basic": {
#                 "name": "cnccU",
#                 "display_name": "人行清算中心U头报文",
#                 "decodes": {
#                     "decode": [
#                         "1827731058",
#                         "1825566008",
#                         "1832061158"
#                     ]
#                 }
#             },
#             "prepare": {
#                 "filed": [
#                     {
#                         "@op": "as",
#                         "@value": "0",
#                         "#text": "ret_code"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "1",
#                         "#text": "is_success"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "120",
#                         "#text": "max_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "20",
#                         "#text": "overtime_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "60",
#                         "#text": "statistics_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "60",
#                         "#text": "delay_span"
#                     },
#                     {
#                         "@op": "as",
#                         "@value": "1",
#                         "#text": "alert_threshold"
#                     },
#                     {
#                         "@op": "as",
#                         "@item1": "CnccOrigReceiverSID",
#                         "#text": "sys_type"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgType",
#                         "#text": "CnccMesgType"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgType",
#                         "#text": "RRB"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgID",
#                         "#text": "MesgId"
#                     },
#                     {
#                         "@op": "str_trim",
#                         "@item1": "CnccMesgRefID",
#                         "#text": "MesgRefId"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "196",
#                         "@end": "216",
#                         "#text": "CCPCId"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "73",
#                         "@end": "88",
#                         "#text": "mbfe_sts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "153",
#                         "@end": "168",
#                         "#text": "ccpc_cts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "CnccReserve",
#                         "@start": "176",
#                         "@end": "191",
#                         "#text": "ccpc_sts"
#                     },
#                     {
#                         "@op": "str_sub",
#                         "@item1": "DestIp",
#                         "@start": "3",
#                         "@end": "6",
#                         "#text": "sub_trans_type"
#                     }
#                 ]
#             },
#             "normalizes": {
#                 "normalize": [
#                     {
#                         "@name": "CnccOrigSender",
#                         "if": [
#                             {
#                                 "@op": "notNull",
#                                 "@item1": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "notNull",
#                                 "@value": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "eq",
#                                 "@item1": "CnccOrigSender",
#                                 "@item2": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             }
#                         ],
#                         "else": {
#                             "filed": [
#                                 {
#                                     "@op": "as",
#                                     "@value": "--"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_add",
#                                     "@item1": "CnccOrigReceiver",
#                                     "@item2": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_sub",
#                                     "@item1": "CnccReserve",
#                                     "@start": "176",
#                                     "@end": "191"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 }
#                             ]
#                         }
#                     },
#                     {
#                         "@name": "CnccOrigSender",
#                         "if": [
#                             {
#                                 "@op": "notNull",
#                                 "@item1": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             },
#                             {
#                                 "@op": "eq",
#                                 "@item1": "CnccOrigSender",
#                                 "@item2": "CnccOrigSender",
#                                 "field": [
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigSender"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_add",
#                                         "@item1": "CnccOrigReceiver",
#                                         "@item2": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_sub",
#                                         "@item1": "CnccReserve",
#                                         "@start": "176",
#                                         "@end": "191"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     },
#                                     {
#                                         "@op": "str_trim",
#                                         "@item1": "CnccOrigReceiver"
#                                     }
#                                 ]
#                             }
#                         ],
#                         "else": {
#                             "filed": [
#                                 {
#                                     "@op": "as",
#                                     "@value": "--"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_add",
#                                     "@item1": "CnccOrigReceiver",
#                                     "@item2": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_sub",
#                                     "@item1": "CnccReserve",
#                                     "@start": "176",
#                                     "@end": "191"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 },
#                                 {
#                                     "@op": "str_trim",
#                                     "@item1": "CnccOrigReceiver"
#                                 }
#                             ]
#                         }
#                     },
#                     {
#                         "@name": "cncc",
#                         "filed": [
#                             {
#                                 "@op": "as",
#                                 "@value": "--"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_add",
#                                 "@item1": "CnccOrigReceiver",
#                                 "@item2": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_sub",
#                                 "@item1": "CnccReserve",
#                                 "@start": "176",
#                                 "@end": "191"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             },
#                             {
#                                 "@op": "str_trim",
#                                 "@item1": "CnccOrigReceiver"
#                             }
#                         ]
#                     }
#                 ]
#             },
#             "traces": {
#                 "filed": [
#                     {
#                         "@name": "critical_keyword",
#                         "field": {
#                             "@op": "as",
#                             "@item1": "transaction_id"
#                         }
#                     },
#                     {
#                         "@name": "fields",
#                         "field": [
#                             {
#                                 "@op": "as",
#                                 "@item1": "status"
#                             },
#                             {
#                                 "@op": "as",
#                                 "@item1": "TxId"
#                             }
#                         ]
#                     }
#                 ]
#             },
#             "trace_disp": {
#                 "filed": [
#                     {
#                         "@op": "disp",
#                         "@item1": "status",
#                         "#text": "业务状态"
#                     },
#                     {
#                         "@op": "disp",
#                         "@item1": "TxId",
#                         "#text": "明细标识号"
#                     }
#                 ]
#             },
#             "dimensions": {
#                 "dimension": [
#                     {
#                         "@name": "trans_type",
#                         "@display_name": "业务类型",
#                         "@csv_file": "trans_type.csv"
#                     },
#                     {
#                         "@name": "sub_trans_type",
#                         "@display_name": "业务类型",
#                         "@csv_file": "sub_trans_type.csv"
#                     }
#                 ]
#             }
#         }
#     }
# }
#
#
# dp = p["protocol"]["dp.xml"]
# for i in dp:
#     with open("cnccU_a.xml", "a+") as f:
#         t = {
#             "protocol": i
#         }
#         f.write(xmltodict.unparse(t, pretty=True)[38:] + "\n")
#
# conf = p["protocol"]["proto.xml"]
# t = {
#     "protocol": conf
# }
# with open("cnccU_a.xml", "a+") as f:
#     f.write(xmltodict.unparse(t, pretty=True)[38:] + "\n")
