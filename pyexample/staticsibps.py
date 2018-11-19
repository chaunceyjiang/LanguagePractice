#!/opt/python27/bin/python

import glob
CCPC_NPC_UP = {'1':0, '2':0}
NPC_DOWN = {'1':0, '2':0}
CCPC_DOWN = {'1':0, '2':0}
LEN = 0
NPC_UP_CCPC = {'saps_01': 0, 'saps_02': 0, 'ccms911': 0, 'ccms900': 0}
for fp in glob.iglob('/opt/cbms/scripts/20180817/convert/*.txt'):
    with open(fp, 'r') as f:
        data = f.read().split('\n')[:-1]
        for i in data:
            dat = eval(i)
            src = dat.get('SrcIp', 0)
            dst_port = dat.get('DestPort', 0)
            dst_ip = dat.get('DestIp', 0)
            src_port = dat.get('SrcPort', 0)
            cncc_type = dat.get('CnccMesgType').strip()
            LEN+=1
            if (dst_ip in [197271584, 197271585, 197271586, 197271587, 197271590, 197271593]) and dst_port == 2531:
                if 'ibps.101.001.01' in cncc_type:
                    CCPC_NPC_UP['1'] += 1
                elif 'ibps.101.001.02' in cncc_type:
                    CCPC_NPC_UP['2'] += 1
                elif 'ibps.102.001.01' in cncc_type:
                    NPC_DOWN['1'] += 1
                elif 'ibps.102.001.02' in cncc_type:
                    NPC_DOWN['2'] += 1
            elif src in [197271584, 197271585, 197271586, 197271587, 197271590, 197271593]:
                if 'ibps.101.001.01' in cncc_type:
                    CCPC_DOWN['1'] += 1
                elif 'ibps.101.001.02' in cncc_type:
                    CCPC_DOWN['2'] += 1
                elif 'saps.601.001.01' in cncc_type:
                    NPC_UP_CCPC['saps_01'] += 1
                elif 'saps.601.001.02' in cncc_type:
                    NPC_UP_CCPC['saps_02'] += 1
                elif 'ccms.900.001' in cncc_type:
                    NPC_UP_CCPC['ccms900'] += 1
                elif 'ccms.911.001' in cncc_type:
                    NPC_UP_CCPC['ccms911'] += 1

with open('/tmp/tongji1', 'w') as f:
    f.write('CCPC_NPC_UP:' + str(CCPC_NPC_UP) + '\n')
    f.write('NPC_DOWN:'+str(NPC_DOWN) + '\n')
    f.write('CCPC_DOWN:'+str(CCPC_DOWN) + '\n')
    f.write('NPC_UP_CCPC:'+str(NPC_UP_CCPC) + '\n')
    f.write('LNE:'+ str(LEN))




