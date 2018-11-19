import socket
import struct
def ipv4_text2int(ip):
    return struct.unpack('!i', socket.inet_aton(ip))[0]


def ipv4_int2text(n):
    return socket.inet_ntoa(struct.pack('!i', n))