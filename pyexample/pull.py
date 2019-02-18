import zmq
import msgpack
import time
ctx = zmq.Context()
socket = ctx.socket(zmq.PUSH)
text = {
    "a":1,
    "b":"asdasdasda",
}
socket.bind("tcp://192.168.2.111:5000")
while True:
    text['ts']=time.time()
    socket.send(msgpack.packb(text))
    print time.time()
    time.sleep(1)
