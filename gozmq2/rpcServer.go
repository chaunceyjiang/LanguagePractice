package main

/* import (
	"context"
	"io"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

// type HelloService struct {
// }

// func (p *HelloService) Hello(request *String, reply *String) error {
// 	// *reply = "hello:" + request
// 	reply.Value = "hello:" + request.GetValue()
// 	return nil
// }

// func main() {
// 	rpc.RegisterName("HelloService", new(HelloService))
// 	listener, err := net.Listen("tcp", ":1234")
// 	if err != nil {
// 		log.Fatal("ListenTCP error", err)
// 	}
// 	conn, err := listener.Accept()
// 	if err != nil {
// 		log.Fatal("Accept error:", err)
// 	}
// 	rpc.ServeConn(conn)

// }
type HelloServiceImpl struct {
}

func (self *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{
		Value: "hello:" + args.GetValue(),
	}
	return reply, nil
}

func (self *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		replay := &String{
			Value: "Server hello:" + args.GetValue(),
		}
		err = stream.Send(replay)
		if err != nil {
			return err
		}
	}
}
func main() {
	log.Println("Runing GRPC...")
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
 */