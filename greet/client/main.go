package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Amartya-Bhardwaj/grpc/greet/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func Dogreet(c pb.GreetServiceClient) {
	log.Println("Dogreet function invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Amartya",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s\n", r.GetResult())
}

func DoGreetMany(c pb.GreetServiceClient) {
	log.Println("DoGreetMany function invoked")
	req := &pb.GreetRequest{
		FirstName: "AMArtya",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
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
		log.Printf("GreetManytimes : %s", msg.GetResult())
	}
}

func DoLongGreet(c pb.GreetServiceClient) {
	log.Println("DoLongGreet invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Amartya"},
		{FirstName: "Harsh"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, req := range reqs {
		log.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result : %s", res.GetResult())
}

func DoManyGreet(c pb.GreetServiceClient) {
	log.Println("DoMany function invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Ahdijds"},
		{FirstName: "slfhifdd"},
		{FirstName: "sfkbfbusdf"},
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
			log.Printf("Result %s", msg.GetResult())
		}
		close(waitc)
	}()
	<-waitc
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	//DoGreetMany(c)
	Dogreet(c)
	//DoLongGreet(c)
	//DoManyGreet(c)
}
