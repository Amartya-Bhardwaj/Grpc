package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func DoSum(c pb.SumServiceClient) {
	log.Println("Do sum invoked")
	req := &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 3,
	}
	r, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sum : %d", r.GetResult())
}

//Server will send many response
func DoPrime(c pb.SumServiceClient) {
	log.Println("DoPrime invoked")
	req := &pb.PrimeRequest{
		Number: 120,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Primes: %d", msg.GetResult())
	}
}

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewSumServiceClient(conn)
	DoSum(c)
	DoPrime(c)
}
