package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Amartya-Bhardwaj/grpc/greet/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

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

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	DoGreetMany(c)
	Dogreet(c)
}
