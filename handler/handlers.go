package handler

import (
	"context"
	"fmt"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedToDoServer
}

func (s *Server) CreateTodoService(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	todo, err := executeInsertStatement(in.Description)
	return &pb.TodoResponse{Id: todo.Id}, err
}

func (s *Server) DeleteTodoService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	return executeDeleteStatement(in.Id)
}

func (s *Server) GetAllTodosService(in *pb.Null, stream grpc.ServerStreamingServer[pb.Todo]) error {
	todos, err := executeSelectStatement("SELECT * FROM todo")
	if err != nil {
		return err
	}
	for i := range todos {
		stream.Send(&todos[i])
	}
	return nil
}

func (s *Server) GetTodoByIDService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	todos, err := executeSelectStatement(fmt.Sprintf("SELECT * FROM todo where id = %s", in.Id))
	if err != nil {
		return nil, err
	}
	return &todos[0], nil
}

func (s *Server) UpdateTodoService(ctx context.Context, in *pb.Todo) (*pb.Todo, error) {
	err := executeUpdateStatement(in)
	if err != nil {
		return nil, err
	}
	return in, nil
}
