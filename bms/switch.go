package main

import (
	. "bms/protocol"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	zmq "github.com/pebbe/zmq4"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"os"
	"path"

)

var (
	protocolpath string
)

func main() {
	var p *ProtocolConfig
	var err error
	protocolpath = "/home/chauncey/Goexample/bms/protocol"
	usage()
	var filename = path.Join(protocolpath, "ibps.xml")
	var body = map[string]interface{}{"ts": 1548927980878719000, "FlowId": "847257491325722656:-7102165741700382719:0", "PktLen": 594, "MetaType": 1, "StreamId": 1, "PartId": 3, "DecodeId": 1832061158, "FlowSide": 0, "SrcIp": 197267507, "DestIp": 197271584, "SrcPort": 40304, "DestPort": 2531, "IpProt": 6, "TcpPldLen": 540, "PktId": 2006919916, "TcpSeq": 2147922445, "TcpAck": 432877654, "CnccOrigReceiverSID": "IBPS", "CnccMesgType": "ibps.101.001.02     ", "CnccMesgID": "10410000000456753317", "CnccMesgRefID": "10410000000456753317", "MsgId": "1041000000042018102653987380", "tranType": "C200", "RRA": "undefine", "Prot": "xml"}
	file, err := os.Open(filename)
	p = NewProtocolConfig(protocolpath)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	xml.Unmarshal(data, &p)
	//fmt.Println("---", p.Prepare.Text, "---")
	p.PrepareHandler(body)
	p.NormalizesHandler(body)
	//fmt.Printf("-%v-", s)
	fmt.Println("-------------------------")
	//fmt.Println(strings.Trim("     x    ", " "), body)
	//xk:=body["mesgtype"]
	fmt.Println(body["trans_type"], body["transaction_id"], body["RRB"], body["MsgId"])
	//fmt.Println(strings.Contains("req,resp","req"))
	//fmt.Println(p.Normalizes.Normalize[0].If)

}

func usage() {
	flag.StringVar(&protocolpath, "protocoldir", protocolpath, "")
	flag.Parsed()
}

func newSoeckt(endpoint string, hwm int) *zmq.Socket {

	ctx, err := zmq.NewContext()
	if err != nil {
		log.Error(err)
	}
	socket, err := ctx.NewSocket(zmq.PULL)
	if err != nil {
		log.Error(err)
	}
	err = socket.SetRcvhwm(hwm)
	if err != nil {
		log.Error(err)
	}
	if err = socket.Connect(endpoint); err != nil {
		log.Error(err)
	}
	return socket
}
func zmqHandler(soc *zmq.Socket, msgch chan<- map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	var msgbody map[string]interface{}
	mh := new(codec.MsgpackHandle)
	mh.RawToString = true
	mh.WriteExt = true
	for {

		parts, err := soc.RecvMessageBytes(0)
		if err != nil {
			log.Error(err)
		}
		// head := parts[0]   头部信息,其实可以不需要,所以这里进行舍弃
		body := parts[1]

		dec := codec.NewDecoderBytes(body, mh)
		err = dec.Decode(&msgbody)
		msgch <- msgbody

	}
}
// TODO 这里只做了协议的转码,但是判断双向交易目前还没有办法做
// protoHanlder  进行协议转码
func protoHanlder(p *ProtocolConfig, bodych <-chan map[string]interface{}, msgch chan<- map[string]interface{}) {
	for body := range bodych {
		go func() {
			p.Process(body)
			msgch <- body
		}()
	}

}
