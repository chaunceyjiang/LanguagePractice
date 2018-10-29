package main
//
//import (
//	"flag"
//	"io/ioutil"
//	"log"
//
//	zmq "github.com/alecthomas/gozmq"
//)
//
//func main() {
//	head_filename := flag.String("headfile", "org_header", "headfile name")
//	body_filename := flag.String("bodyfile", "org_body", "bodyfile name")
//	sendNum := flag.Int64("sendnum", 10, "send number")
//	//goNum := flag.Int("gonum", 10, "goruntine number")
//	flag.Parse()
//	context, err := zmq.NewContext()
//	if err != nil {
//		log.Fatal(err)
//	}
//	socket, err := context.NewSocket(zmq.PUSH)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = socket.Connect("tcp://192.168.1.253:23220")
//	if err != nil {
//		log.Fatal(err)
//	}
//	bodys,err:=GetDate(body_filename)
//	heads,err:=GetDate(head_filename)
//	for j := int64(0); j < *sendNum; j++ {
//		err = socket.Send(bodys, zmq.SNDMORE)
//		if err != nil {
//			log.Fatal(err)
//		}
//		err = socket.Send(heads, zmq.SNDMORE)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
//func GetDate(filename *string) ([]byte, error) {
//	body, err := ioutil.ReadFile(*filename)
//	log.Printf("%d\n", len(body))
//	return body, err
//}
