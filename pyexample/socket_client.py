# coding: utf-8
import json
import struct
import socket
import time
def rpc(sock,in_,params):
    request = json.dumps({"in_":in_,"params":params})
    length_prefix = struct.pack("I",len(request)) #"I" unsigned int  4 bytes
    sock.sendall(length_prefix)#报文长度
    sock.sendall(request)
    length_prefix = sock.recv(4)
    length = struct.unpack('I',length_prefix)
    body = sock.recv(length)
    response = json.loads(body)
    return response["out"],response["result"]

if __name__ == "__main__":
    s = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    s.connetc(("localhost",8000))
    for i in range(10):  # 连续发送10个rpc请求
        out, result = rpc(s, "ping", "ireader %d" % i)
        print out, result
        time.sleep(1)  # 休眠1s，便于观察
    s.close()  # 关闭连接