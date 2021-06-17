package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jimmyjimmbles/first-api/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting Unary RPC")
	req := &calcpb.IntegersRequest{
		Integers: &calcpb.Integers{
			Int1: 10,
			Int2: 30,
		},
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Greet RPC: %v", err)
	}

	log.Printf("Res from Greet %v", res.Sum)
}
