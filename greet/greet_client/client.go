package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jimmyjimmbles/first-api/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Print("Hello I am Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jay",
			LastName: "Jones",
		},
	}	
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Greet RPC: %v", err)
	}

	log.Printf("Res from Greet %v", res.Result)
}
