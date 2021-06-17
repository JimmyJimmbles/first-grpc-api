package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jimmyjimmbles/first-api/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calcpb.IntegersRequest) (*calcpb.IntegersResponse, error) {
	fmt.Printf("Calculator Function Invoked %v\n", req)
	firstNum := req.GetIntegers().GetInt1()
	secondNum := req.GetIntegers().GetInt2()
	sum := firstNum + secondNum
	res := &calcpb.IntegersResponse{
		Sum: sum,
	}

	return res, nil
}

func main() {
	fmt.Println("Server Started")

	// Listen to port
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	// Init and register grpc server
	s := grpc.NewServer()
	// calcpb.RegisterCalculatorServiceServer(s, &server{})
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
