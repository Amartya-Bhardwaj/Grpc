package main

import (
	"context"
	"log"

	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function invoked: %v", in)
	return &pb.SumResponse{
		Result: in.GetFirstNumber() + in.GetSecondNumber()}, nil
}
