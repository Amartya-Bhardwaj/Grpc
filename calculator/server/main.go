package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Recieved %v %v", in.GetFirstNumber(), in.GetSecondNumber())
	return &pb.SumResponse{Result: in.GetFirstNumber() + in.GetSecondNumber()}, nil
}

func (s *server) Primes(in *pb.PrimeRequest, stream pb.SumService_PrimesServer) error {
	log.Printf("Recieved %v", in.GetNumber())
	k := (int64)(2)
	number := in.GetNumber()
	for number > 1 {
		if number%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			number = number / k
		} else {
			k = k + 1
		}
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
	pb.RegisterSumServiceServer(s, &server{})
	log.Printf("Listening on %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
