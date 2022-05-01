package main

// import (
// 	"context"
// 	"log"

// 	pb "github.com/Amartya-Bhardwaj/grpc/calculator/proto"
// )

// func doSum(c pb.SumServiceClient) {
// 	log.Printf("DoSum invoked")
// 	res, err := c.Sum(context.Background(), &pb.SumRequest{
// 		FirstNumber:  1,
// 		SecondNumber: 2})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Result: %d", res.Result)
// }
