#!/usr/bin/python
#coding:utf-8

with open('body', 'r') as f:
    datas = f.readlines()

# npc与hvps的IP配置
NpcIp = "11.194.32.32,11.194.32.33,11.194.32.34,11.194.32.35,11.194.32.38,11.194.32.41"
HvpsIp = "11.194.16.16,11.194.16.17,11.194.16.18,11.194.16.19"

# 循环每一行
# Npc到HVPS请求总数
NreqAll = 0
# Npc到HVPS响应总数
NrespAll = 0
# HVPS到NPC请求总数
HreqAll = 0
# HVPS到NPC响应总数

NpcToHvps = []
HvpsToNpc = []
HrespAll = 0
for data in datas:
    dat = data.split('\n')[0]
    dat1 = eval(dat)
    SrcIp = str(dat1.get("SrcIp",0))
    DstIp = str(dat1.get("DestIp",0))
    RRA = str(dat1.get("RRA",0))
    
    # Npc-->HVPS  req
    if (SrcIp in NpcIp) and (DstIp in HvpsIp) and (RRA == "req"):
        NreqAll += 1
        NpcToHvps.insert(len(NpcToHvps), data)
    # Npc-->HVPS resp
    if SrcIp in NpcIp and DstIp in HvpsIp and RRA == "resp":
        NrespAll += 1
        NpcToHvps.insert(len(NpcToHvps), data)
    # HVPS-->NPC req
    if SrcIp in HvpsIp and DstIp in NpcIp and RRA == "req":
        HreqAll += 1
        HvpsToNpc.insert(len(HvpsToNpc), data)
    # HVPS-->NPC resp
    if SrcIp in HvpsIp and DstIp in NpcIp and RRA == "resp":
        HrespAll += 1
        HvpsToNpc.insert(len(HvpsToNpc), data)

print("总交易量统计数据:")
print("N-->H  req: s%", NreqAll)
print("N-->H  resp: s%", NrespAll)
print("H-->N  req: s%", HreqAll)
print("H-->N  resp: s%", HrespAll)


# 统计HVPS的响应总笔数
# 请求数据
respCount = 0
for dat in NpcToHvps:
    data1 = eval(dat)
    MsgId = str(data1.get("transaction_id",0))
    RRA = str(data1.get("RRA", 0))
    # NPC到HVPS的请求
    if RRA == "req":
        # HVPS到NPC的响应
        for dat2 in HvpsToNpc:
            data2 = eval(dat2)
            OrigMsgId = str(data2.get("transaction_id", 0))
            RRA2 = str(data2.get("RRA", 0))
            # 是响应数据
            if RRA2 == "resp":
                # 是同一笔交易
                if OrigMsgId == MsgId:
                    respCount += 1

print("Success resp：s%", respCount)