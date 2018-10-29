package main

/* import (
	"log"

	zmq "github.com/alecthomas/gozmq"
	"github.com/vmihailenco/msgpack"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}

func main() {
	context, err := zmq.NewContext()
	checkErr(err)
	socket, err := context.NewSocket(zmq.PULL)
	err = socket.Bind("tcp://*:23220")
	err = socket.Bind("tcp://*:23221")
	err = socket.Bind("tcp://*:23222")
	err = socket.Bind("tcp://*:23223")
	err = socket.Bind("tcp://*:23224")
	err = socket.Bind("tcp://*:23225")
	err = socket.Bind("tcp://*:23226")
	err = socket.Bind("tcp://*:23227")
	checkErr(err)
	var count int
	var head, body []interface{}
	for {
		parts, _ := socket.RecvMultipart(zmq.NOBLOCK)
		if len(parts) != 0 {
			count ++
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
 */