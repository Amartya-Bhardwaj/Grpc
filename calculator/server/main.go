package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
)

var addr string = "0.0.0.0:8000"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on: %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
