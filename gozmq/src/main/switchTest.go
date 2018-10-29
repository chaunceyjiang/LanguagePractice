package main

//import (
//	"flag"
//	zmq "github.com/alecthomas/gozmq"
//	"io/ioutil"
//	"log"
//	"sync"
//)
//
//func main() {
//	head_filename := flag.String("headfile", "org_header", "headfile name")
//	body_filename := flag.String("bodyfile", "org_body", "bodyfile name")
//	sendNum := flag.Int64("sendnum", 10000, "send number")
//	goNum := flag.Int("gonum", 10, "goruntine number")
//	flag.Parse()
//	context, err := zmq.NewContext()
//	if err != nil {
//		log.Fatal(err)
//	}
//	var sendNums int64
//	var goNums int
//	var body, head []byte
//	goNums = *goNum
//	sendNums = *sendNum
//	wg := sync.WaitGroup{}
//	head, err = GetDate(head_filename)
//	body, err = GetDate(body_filename)
//	for i := 0; i < goNums; i++ {
//		wg.Add(1)
//		go func(wgs *sync.WaitGroup, sendNumsx int64, bodys, heads []byte) {
//
//			defer wgs.Done()
//			socket, err := context.NewSocket(zmq.PUSH)
//			defer socket.Close()
//			if err != nil {
//				log.Fatal(err)
//			}
//			err = socket.Connect("tcp://127.0.0.1:23200")
//			if err != nil {
//				log.Fatal(err)
//			}
//			err = socket.Connect("tcp://127.0.0.1:23201")
//			if err != nil {
//				log.Fatal(err)
//			}
//			err = socket.Connect("tcp://127.0.0.1:23202")
//			if err != nil {
//				log.Fatal(err)
//			}
//			err = socket.Connect("tcp://127.0.0.1:23203")
//			if err != nil {
//				log.Fatal(err)
//			}
//			for j := int64(0); j < sendNumsx; j++ {
//				err = socket.Send(bodys, zmq.SNDMORE)
//				if err != nil {
//					log.Fatal(err)
//				}
//				err = socket.Send(heads, zmq.SNDMORE)
//				if err != nil {
//					log.Fatal(err)


//				}
//			}
//			log.Println("success")
//		}(&wg, sendNums, body, head)
//	}
//	wg.Wait()
//
//	context.Close()
//}
//func GetDate(filename *string) ([]byte, error) {
//	body, err := ioutil.ReadFile(*filename)
//	log.Printf("%d\n", len(body))
//	return body, err
//}

