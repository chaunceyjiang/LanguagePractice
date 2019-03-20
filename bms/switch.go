package main

import (
	. "bms/protocol"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var p *ProtocolConfig
	var err error
	var protocolpath  = "/home/chauncey/Goexample/bms/protocol"
	var filename = path.Join(protocolpath,"ibps.xml")
	var body = map[string]interface{}{"ts": 1548927980878719000, "FlowId": "847257491325722656:-7102165741700382719:0", "PktLen": 594, "MetaType": 1, "StreamId": 1, "PartId": 3, "DecodeId": 1832061158, "FlowSide": 0, "SrcIp": 197267507, "DestIp": 197271584, "SrcPort": 40304, "DestPort": 2531, "IpProt": 6, "TcpPldLen": 540, "PktId": 2006919916, "TcpSeq": 2147922445, "TcpAck": 432877654, "CnccOrigReceiverSID": "IBPS", "CnccMesgType": "ibps.101.001.02     ", "CnccMesgID": "10410000000456753317", "CnccMesgRefID": "10410000000456753317", "MsgId": "1041000000042018102653987380", "tranType": "C200", "RRA": "undefine", "Prot": "xml"}
	file, err := os.Open(filename)
	p = NewProtocolConfig(protocolpath)
	if err != nil {
		log.Println(err)
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
	fmt.Println(body["trans_type"],body["transaction_id"],body["RRB"],body["MsgId"])
	//fmt.Println(strings.Contains("req,resp","req"))
	//fmt.Println(p.Normalizes.Normalize[0].If)
}
