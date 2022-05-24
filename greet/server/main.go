package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Amartya-Bhardwaj/grpc/greet/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Recieved: %v", in.GetFirstName())
	return &pb.GreetResponse{Result: "Hello" + in.GetFirstName()}, nil
}

func (s *server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Recieved: %v", in.GetFirstName())
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, no: %d", in.GetFirstName(), i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

func (s *server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")
	res := ""
	for {
		req, err := stream.Recv()
		log.Printf("Recieve : %s", req.GetFirstName())
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		res += fmt.Sprintf("Hello %s\n", req.GetFirstName())
	}
}

func (s *server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		res := "Hello" + req.GetFirstName() + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	// opts := []grpc.ServerOption{}
	// tls := true
	// if tls {
	// 	certFile := "ssl/server.csr"
	// 	keyFile := "ssl/server.pem"
	// 	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	opts = append(opts, grpc.Creds(creds))
	// }
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	log.Printf("listening on %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
