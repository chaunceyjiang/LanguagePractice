# coding: utf-8
import json
import struct
import os
import time
def send_message(conn,out,result):
    response = json.dumps({"out":out,"result":result})
    length_prefix = struct.pack("I",len(response))
    conn.sendall(length_prefix)
    conn.sendall(response)

def loog(sock,handlers):
    while True:
        conn,addr = sock.accept()
        pid = os.fork()# 创建子进程
        if pid < 0 :
            return #创建失败
        elif pid == 0:
            # 当前处于子进程
            pass
        else:
            pass

i = 1541905810
while True:
    
    print i%16
    i +=60

    time.sleep(1)
