package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type sever struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	s := grpc.NewServer()
}