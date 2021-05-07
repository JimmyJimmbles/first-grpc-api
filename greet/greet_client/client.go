package main

import (
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
	fmt.Printf("Created Client: %f", c)
}