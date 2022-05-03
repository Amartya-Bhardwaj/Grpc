package main

import (
	"context"
	"flag"
	"fmt"
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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	log.Printf("listening on %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
