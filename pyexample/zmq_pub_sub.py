# coding: utf-8
import zmq
import sys
import msgpack
import time

context = zmq.Context()
topicfilter = "10001"

def sub(ip_port_list):
    socket = context.socket(zmq.SUB)
    socket.setsockopt(zmq.SUBSCRIBE, topicfilter)
    print "订阅 topic{}".format(topicfilter)
    for ip_port in ip_port_list:
        print "connect {}".format(ip_port)
        socket.connect(ip_port)
    while True:
        part = socket.recv_multipart()
        if len(part) != 2:
            print "参数错误"
            continue
        print "part[0]: {},part[1]: {}".format(part[0],msgpack.unpackb(part[1]))
        time.sleep(1)

def pub(ip_port_list):
    socket = context.socket(zmq.PUB)
    for ip_port in ip_port_list:
        socket.bind(ip_port)
    while True:
        socket.send_multipart([topicfilter,msgpack.packb(str(time.time()))])
        time.sleep(1)


if __name__== "__main__":
    print sys.argv
    ip_port_list = sys.argv[2:]
    if sys.argv[1] == "sub":
        sub(ip_port_list)
    else:
        pub(ip_port_list)
print "???"