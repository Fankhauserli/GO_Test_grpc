package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedToDoServer
}

func newServer(bindAddress int) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", bindAddress))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to listen on port %d: %v", bindAddress, err))
	}

	s := grpc.NewServer()
	pb.RegisterToDoServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to serve: %v", err))
	}
	return s, nil
}

func (s *server) createTodoService(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	todo, err := executeInsertStatement(in.Description)
	return &pb.TodoResponse{Id: todo.Id}, err
}

func (s *server) deleteTodoService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	return executeDeleteStatement(int(in.Id))
}

func (s *server) getAllTodosService(in *pb.Null, stream grpc.ServerStreamingServer[pb.Todo]) error {
	todos, err := executeSelectStatement("SELECT * FROM todo")
	if err != nil {
		return err
	}
	for _, todo := range todos {
		stream.Send(&todo)
	}
	return nil
}

func (s *server) getTodoByIDService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	todos, err := executeSelectStatement(fmt.Sprintf("SELECT * FROM todo where id = %d", in.Id))
	if err != nil {
		return nil, err
	}
	return &todos[0], nil
}

func (s *server) updateTodoService(ctx context.Context, in *pb.Todo) (*pb.Todo, error) {
	err := executeUpdateStatement(in)
	if err != nil {
		return nil, err
	}
	return in, nil
}
