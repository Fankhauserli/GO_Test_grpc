package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Fankhauserli/GO_Test_grpc/handler"
	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 20002))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to listen on port %d: %v", 20002, err))
	}

	s := grpc.NewServer()
	pb.RegisterToDoServer(s, &handler.Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(fmt.Errorf("failed to serve: %v", err))
	}
}
