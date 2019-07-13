package main

/* import (
	"encoding/binary"
	"fmt"
)

var IPTCP = []byte{
	0x45, 0x00, 0x00, 0x8c, 0x28, 0xd1, 0x00, 0x00, 0xff, 0x06, 0x00, 0x00, 0x73, 0xef, 0xd2, 0x1b,
	0xac, 0x15, 0x00, 0x01, 0x00, 0x50, 0xe7, 0xa3, 0x93, 0x2d, 0xac, 0xdb, 0x9d, 0x0e, 0x0f, 0x41,
	0x50, 0x10, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x34, 0x70, 0x78, 0x3b, 0x70, 0x61, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x31, 0x30, 0x70, 0x78, 0x3b, 0x70, 0x61,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x31, 0x30, 0x70, 0x78,
	0x3b, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x3a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3b,
	0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x6e, 0x6f, 0x77, 0x72,
	0x61, 0x70, 0x7d, 0x2e, 0x63, 0x2d, 0x64, 0x72, 0x6f, 0x70, 0x64, 0x6f,
}

var IP = []byte{
	0x45, 0x00, 0x02, 0x87, 0xa0, 0x06, 0x40, 0x00, 0x3c, 0x06, 0x00, 0x00, 0x0a, 0x01,
	0x22, 0x15, 0x0a, 0x01, 0x21, 0x0f,
}

// var IP []byte = []byte{0x45, 0x00, 0x00, 0x30, 0x80, 0x4c, 0x40, 0x00, 0x80, 0x06, 0x00, 0x00, 0xd3, 0x43, 0x11, 0x7b, 0xcb, 0x51, 0x15, 0x3d}
var IP14 = []byte{
	0x45 ,0x00 ,0x02 ,0x87 ,0xa0 ,0x06 ,0x40 ,0x00 ,0x3c ,0x06 ,0x00 ,0x00 ,0x0a ,0x01,
0x22 ,0x15 ,0x0a ,0x01 ,0x21 ,0x0f ,0xb7 ,0x18 ,0x9c ,0x45 ,0xc6 ,0x1d ,0x91 ,0x10 ,0x6f ,0x1d,
0x4b ,0x58 ,0x50 ,0x18 ,0xff ,0xff ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x02 ,0x5b ,0x02 ,0x08,
0x00 ,0x00 ,0x0a ,0x41 ,0x41 ,0x41 ,0x41 ,0x41 ,0x41 ,0x41 ,0x41 ,0x58 ,0x43 ,0x31 ,0x31 ,0x73,
0x77 ,0x69 ,0x74 ,0x63 ,0x68 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x63 ,0x75 ,0x70 ,0x73 ,0x70 ,0x69,
0x6b ,0x00 ,0x00 ,0x00 ,0x00 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30,
0x00 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x00 ,0x50 ,0x49 ,0x4b,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x31 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x31 ,0x31 ,0x73 ,0x77,
0x69 ,0x74 ,0x63 ,0x68 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x76 ,0x62 ,0x73 ,0x70 ,0x69 ,0x6b ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x00,
0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x30 ,0x00 ,0x50 ,0x49 ,0x4b ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x31 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00,
0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x00 ,0x2e ,0x31 ,0x32 ,0x36 ,0x45 ,0x42,
0x35 ,0x34 ,0x33 ,0x35 ,0x43 ,0x46 ,0x38 ,0x45 ,0x42 ,0x31 ,0x43 ,0x39 ,0x46 ,0x30 ,0x31 ,0x30,
0x31 ,0x30 ,0x30 ,0x38 ,0x32 ,0x30 ,0x32 ,0x38 ,0x32 ,0x39 ,0x32 ,0x31 ,0x37 ,0x30 ,0x30 ,0x38,
0x32 ,0x30 ,0x32 ,0x38 ,0x32 ,0x39 ,0x32 ,0x31 ,0x37,
}
func checksum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length == 1 {
		sum += uint32(data[index]) << 8
	}
	sum += (sum >> 16)
	return ^uint16(sum)
}
func getPseudoHeader(data []byte) []byte { //12 bytes
	var pseudoHeader []byte
	sourceIP := data[12:16]
	pseudoHeader = append(pseudoHeader, sourceIP...)
	destinationIP := data[16:20]
	pseudoHeader = append(pseudoHeader, destinationIP...)
	zeros := []byte{0x00}
	pseudoHeader = append(pseudoHeader, zeros...)
	protocol := data[9]
	pseudoHeader = append(pseudoHeader, protocol)
	totalLen := binary.BigEndian.Uint16(data[2:4])
	ipHeaderLen := (uint16(data[0]) & 0x000f) * 4
	tcpLen := totalLen - ipHeaderLen
	tcpLength := make([]byte, 2)
	binary.BigEndian.PutUint16(tcpLength, tcpLen)
	pseudoHeader = append(pseudoHeader, tcpLength...)
	return pseudoHeader
}

func main() {
	fmt.Printf("%x\n", checksum(IP14[:20]))
	pseudoHeader := append(getPseudoHeader(IP14), IP14[20:]...)
	fmt.Printf("%x\n", checksum(pseudoHeader))

}
*/
