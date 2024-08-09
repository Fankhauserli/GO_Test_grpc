package main

import (
	"context"
	"log"
	"net"

	pb "TodoGrpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, World! "}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":54321")
	if err != nil {
		log.Fatalf("failed to listen on port 54321: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
