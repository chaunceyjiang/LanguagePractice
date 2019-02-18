package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	bashPath     = "/home/chauncey/Desktop/shm"
	recvFilename = bashPath + "/" + "eth1_cmd_recv"
	sendFilename = bashPath + "/" + "eth1_cmd_send"
	dataFilename = bashPath + "/" + "eth1_data"
)

func main() {
	recv_data, _ := ioutil.ReadFile(recvFilename)
	send_data, _ := ioutil.ReadFile(sendFilename)
	//data, _ := ioutil.ReadFile(dataFilename)

	bytesBuffer := bytes.NewBuffer(recv_data)
	fmt.Println(send_data)
	fmt.Println(recv_data)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	file,_:=os.Open(dataFilename)
	file.Seek(int64(x), io.SeekCurrent)
	var data []byte
	file.Read(data)
	fmt.Println(data)
	fmt.Println(x)

}
