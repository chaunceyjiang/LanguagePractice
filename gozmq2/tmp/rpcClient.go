package main

/*
import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	grpc "google.golang.org/grpc"
)


// func main() {
// 	client,err:=rpc.Dial("tcp", ":1234")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var reply string
// 	err = client.Call("HelloService.Hello", "hellox",&reply)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(reply)
// }


func main() {
	fmt.Println("Connetion GRPC Server...")
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{
		Value: "hello grpcClient",
	})
	if err != nil {
		log.Fatal(err)
	}
	streamReply, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := streamReply.Send(&String{
				Value: "client hi",
			}); err != nil {
				time.Sleep(time.Second * 1)
			}
		}
	}()

	for {
		replyS, err := streamReply.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(replyS.GetValue())
	}
	fmt.Println(reply.GetValue())
}
*/
