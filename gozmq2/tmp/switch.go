package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/ugorji/go/codec"
)

var logs *log.Logger

var ticker = time.NewTicker(time.Second * 10)
func init() {
	logs = log.New(os.Stdout, "", log.LUTC|log.Ltime|log.Lshortfile)
	go func() {
		for t := range ticker.C{
			logs.Println("============now====================", time.Now().Unix())
		}
	}()
}

// var orgHead = []int{1526284784, 70391000, 1618494596, 0, 67410548, 0}
func main() {
	ctx, _ := zmq.NewContext()
	recvSocket, _ := ctx.NewSocket(zmq.PULL)
	var diffValue int64 = 1
	if len(os.Args) <= 1 {
		fmt.Printf("%s PORT TIMEDIFF", os.Args[0])
		os.Exit(1)
	}
	endpoint := fmt.Sprintf("tcp://*:%s", os.Args[1])
	if len(os.Args) > 2 {
		diffValue, _ = strconv.ParseInt(fmt.Sprintf("%s", os.Args[2]), 10, 64)
	}
	fmt.Println(endpoint)
	recvSocket.Bind(endpoint)
	recvSocket.SetHWM(100000)
	// var head, body []byte
	var head []byte
	var msghead []interface{}
	// var msgbody map[string]interface{}
	ticker2 := time.NewTicker(time.Second * diffValue)

	go func(){
		for t := range ticker2.C{
			fmt.Printf("%d head: %v ", t, msghead)
		}
	}()
	for {
		parts, err := recvSocket.RecvMultipart(0)
		if err != nil {
			logs.Fatalln(err)
		}
		head = parts[0]
		mh := new(codec.MsgpackHandle)
		mh.RawToString = true
		dec := codec.NewDecoderBytes(head, mh)
		dec.Decode(&msghead)
		t := time.Now().Unix()
		// if uint64(t)-msghead[0].(uint64) >= uint64(diffValue) {
		// 	fmt.Printf("%d head: %v ", t, msghead)
		// 	fmt.Printf(" diff: %d \n", uint64(t)-msghead[0].(uint64))
		// }
		// body = parts[1]
		// dec = codec.NewDecoderBytes(body, mh)
		// dec.Decode(&msgbody)
		// fmt.Printf("%d body: %v\n", t, msgbody)

	}
}
