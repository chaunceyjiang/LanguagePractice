import zmq
import msgpack
import time
ctx = zmq.Context()

class Recv(object):

    def __init__(self):
        self.socket = ctx.socket(zmq.PULL)
        self.socket.bind("tcp://192.168.1.24:23200")
    def recv(self):
        return self.socket.recv_multipart()

class Send(object):

    def __init__(self):
        self.socket = ctx.socket(zmq.PUSH)
        self.socket.bind("tcp://192.168.1.24:5555")
    def send(self,parts):
        self.socket.send_multipart(parts)

if __name__=="__main__":
    r = Recv()
    s = Send()
    while True:
        s.send(r.recv())
