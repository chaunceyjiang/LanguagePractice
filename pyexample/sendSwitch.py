# coding: utf-8
import msgpack,zmq,os,math,random,time
endpoint = "tcp://192.168.2.111:5000"

PATH = "/home/chauncey/Desktop/tmp/"
filePackList = []
def listPack():
    for filename in os.listdir(PATH):
        if str(filename).endswith(".pack"):
            filePackList.append(PATH+filename)
def unpack_raw(stream):
    unpacker = msgpack.Unpacker(stream)
    for msg in unpacker:
        yield msg

def unpack_msg(msg):
    return msgpack.unpackb(msg)

def pack_msg(pyobj):
    return msgpack.packb(pyobj)

def SendSwitch():
    context = zmq.Context()
    socket = context.socket(zmq.PUSH)
    socket.bind(endpoint)
    while True:
        for filename in filePackList:
            # print filename
            for body in unpack_raw(file(filename)):
                if 'DP#SYNC' not in body:
                    head = [None] * 6
                    head[0] = body['ts'] / 1000000000 + 57600 + 18000
                    head[1] = int( round(body['ts'] / 1000000000.0 -body['ts'] / 1000000000,9)  * 1000000000)
                    # print head[0],head[1]
                    body['ts'] = head[0] * 1000000000 + head[1]
                    head[2] = body['DecodeId']
                    head[3] = random.randint(1,6)
                    head[4] = head[2] 
                    head[5] = 0
                    if head[3] == 3:
                        print body
                    socket.send_multipart([pack_msg(head),pack_msg(body)])
listPack()
SendSwitch()
time.sleep(20)