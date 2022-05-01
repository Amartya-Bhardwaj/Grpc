package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
)

var addr string = "localhost:8000"

func doSum(c pb.SumServiceClient) {
	log.Printf("DoSum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result: %d", res.Result)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewSumServiceClient(conn)
	doSum(c)
}
