package main
/* 
import (
	"bytes"
	"log"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/ugorji/go/codec"
)

var orgHead = []int{1526538113, 343541000, 1618494596, 0, 67410548, 0}

var orgBody = map[string]interface{}{"ts": 1526538113343541000,
	"FlowId": "720894889111978263:-7633500475215577087:0",
	"PktLen": 361, "MetaType": 1, "StreamId": 1, "PartId": 0,
	"DecodeId": 1618494596, "Vlan": 34, "FlowSide": 0, "SrcIp": 167846421,
	"DestIp": 166330647, "SrcPort": 38416, "DestPort": 23491, "IpProt": 6, "TcpPldLen": 303,
	"PktId": 2241, "TcpSeq": 1851467491, "TcpAck": 1349491582, "MsgType": "0210",
	"PAN": "6210-XXXX-XXXX-8747", "PanHash": "4f6d81bf694a871a22c48ceac64d0a01",
	"F3": "000000", "F7": "0521175153", "F11": "606207", "F25": "00", "F32": "48028810",
	"RespCode": "00", "Lookup": "02100000003", "TransId": "0521175153:606207:48028810:00098800",
	"RRA": "resp", "Prot": "cups"}

func main() {
	ctx, _ := zmq.NewContext()
	sendSocket, err := ctx.NewSocket(zmq.PUSH)
	sendSocket.SetHWM(100000)
	if err != nil {
		log.Fatalln(err)
	}
	var msg [][]byte
	var record []byte

	w := new(bytes.Buffer)
	mh := new(codec.MsgpackHandle)
	enc := codec.NewEncoder(w, mh)
	enc.Encode(orgHead)
	record = w.Bytes()
	msg = append(msg, record)

	w = new(bytes.Buffer)
	enc = codec.NewEncoder(w, mh)
	enc.Encode(orgBody)
	record = w.Bytes()
	msg = append(msg, record)
	err = sendSocket.Connect("tcp://192.168.1.112:23200")
	// err = sendSocket.Connect("tcp://*:23200")
	// err = sendSocket.Connect("tcp://*:23203")
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 4; i++ {
		sendSocket.SendMultipart(msg, 0)
		// sendSocket.Send(msg[0], zmq.SNDMORE)
		// sendSocket.Send(msg[1], 0)

		time.Sleep(time.Second * 2)
	}
	// sendSocket.SendMultipart(msg, 0)
}
 */