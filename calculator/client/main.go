package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

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

func DoAverage(c pb.SumServiceClient) {
	log.Println("DoAverage invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, req := range reqs {
		log.Printf("Sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Avg of the requested number : %0.2f", res.GetResult())
}

func DoMain(c pb.SumServiceClient) {
	log.Println("Function DoMain invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Max number: %d", msg.GetResult())
		}
		close(waitc)
	}()
	<-waitc
}

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewSumServiceClient(conn)
	// DoSum(c)
	//DoPrime(c)
	DoAverage(c)
	//DoMain(c)
}
