
import msgpack
import sys

def unpack_raw(stream):
    unpacker = msgpack.Unpacker(stream)
    for msg in unpacker:
        yield msg


def unpack_msg(msg):
    return msgpack.unpackb(msg)


def pack_msg(pyobj):
    return msgpack.packb(pyobj)


if __name__ == '__main__':
    fd = file(sys.argv[2],'w')
    for msg in unpack_raw(open(sys.argv[1])):
    #for msg in unpack_msg(sys.argv[1]):
        fd.write(str(msg))
        fd.write('\n')
    fd.close()
