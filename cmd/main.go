package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedToDoServer
	db *sqlx.DB
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", bindAddress))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to listen on port %d: %v", bindAddress, err))
	}

	s := grpc.NewServer()
	var db = newDB()

	pb.RegisterToDoServer(s, &server{db: db})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

}

func newServer(bindAddress int) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", bindAddress))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to listen on port %d: %v", bindAddress, err))
	}

	s := grpc.NewServer()
	var db = newDB()

	pb.RegisterToDoServer(s, &server{db: db})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return s, nil
}
