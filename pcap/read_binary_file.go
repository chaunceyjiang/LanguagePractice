package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	bashPath = "/home/chauncey/Desktop/shm"
	recvFilename     = bashPath + "/" + "eth1_cmd_recv"
	sendFilename     = bashPath + "/" + "eth1_cmd_send"
	dataFilename     = bashPath + "/" + "eth1_data"
)

func main() {
	recv_data, _ := ioutil.ReadFile(recvFilename)
	send_data, _ := ioutil.ReadFile(sendFilename)
	data,_:=os.Open(dataFilename)
	second:=make([]byte,4)
	y:=binary.BigEndian.Uint32(send_data[:4])
	buf:=bufio.NewReader(data)
	
	buf.Discard(int(y))
	buf.Read(second)
	fmt.Println(second)
	fmt.Println(buf.Size())
	fmt.Println(recv_data)
	fmt.Println(send_data)

	i:=binary.BigEndian.Uint32(recv_data[:4])

	t:=binary.BigEndian.Uint32(second)
	fmt.Println(i)
	fmt.Println(y)
	fmt.Println(t)


}
