package main

import (
	"log"
	zmq "github.com/alecthomas/gozmq"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

func main() {
	context, err := zmq.NewContext()
	checkErr(err)
	socket, err := context.NewSocket(zmq.PULL)
	err = socket.Bind("tcp://*:23200")
	err = socket.Bind("tcp://*:23201")
	err = socket.Bind("tcp://*:23202")
	err = socket.Bind("tcp://*:23203")
	err = socket.Bind("tcp://*:23204")
	err = socket.Bind("tcp://*:23205")
	err = socket.Bind("tcp://*:23206")
	err = socket.Bind("tcp://*:23207")
	checkErr(err)
	var count int
	var head, body []interface{}
	for {
		parts, _ := socket.RecvMultipart(zmq.NOBLOCK)
		if len(parts) != 0 {
			count++
			msgpack.Unmarshal(parts[0], &head)
			msgpack.Unmarshal(parts[1], &body)
			log.Printf("%v \n %v \n", head, body)

			log.Println(count)
		}

	}
}
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}

}


 