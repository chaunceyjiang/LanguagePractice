package main

// func main() {
// 	records := [][]string{
// 		{"first_name", "last_name", "username"},
// 		{"Rob", "Pike", "rob"},
// 		{"Ken", "Thompson", "ken"},
// 		{"Robert", "Griesemer", "gri"},
// 	}
// 	f,err:=os.Create("channel.csv")
// 	f.Seek(1, whence 0)
// 	if err!=nil{
// 		log.Fatalln(err)
// 	}
// 	w := csv.NewWriter(f)

// 	for _, record := range records {
// 		if err := w.Write(record); err != nil {
// 			log.Fatalln("error writing record to csv:", err)
// 		}
// 	}

// 	// Write any buffered data to the underlying writer (standard output).
// 	w.Flush()

// 	if err := w.Error(); err != nil {
// 		log.Fatal(err)
// 	}
// }

/* import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
)

func main()  {
	f,err:=os.Open("trans_type.csv")
	if err!=nil{
		log.Fatalln(err)
	}
	f.Seek(15, 0)
	r:=csv.NewReader(f)
	record,_:=r.Read()
	fmt.Println(record)
} */

/* import (
	"fmt"

	zmq "github.com/alecthomas/gozmq"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func main() {

	ctx, err := zmq.NewContext()
	if err != nil {
		fmt.Println(err)
	}
	zmqsocket, err := ctx.NewSocket(zmq.PULL)
	if err != nil {
		fmt.Println(err)
	}
	err = zmqsocket.Connect("tcp://192.168.1.113:5561")
	if err != nil {
		fmt.Println(err)
	}
	// x,_:=msgpack.Marshal("hello world")
	// err = zmqsocket.Send(x, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	for i := 0; i < 10; i++ {

		data, err := zmqsocket.Recv(0)
		if err != nil {
			fmt.Println(err)
		}
		var s string
		msgpack.Unmarshal(data, &s)
		fmt.Println(string(data))
	}

}
 */