package main

import (
	"context"
	"fmt"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"google.golang.org/grpc"
)

func (s *models.server) CreateTodoService(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	todo, err := executeInsertStatement(in.Description, s.db)
	return &pb.TodoResponse{Id: todo.Id}, err
}

func (s *models.server) DeleteTodoService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	return executeDeleteStatement(int(in.Id), s.db)
}

func (s *models.server) GetAllTodosService(in *pb.Null, stream grpc.ServerStreamingServer[pb.Todo]) error {
	todos, err := executeSelectStatement("SELECT * FROM todo", s.db)
	if err != nil {
		return err
	}
	for _, todo := range todos {
		stream.Send(&todo)
	}
	return nil
}

func (s *models.server) GetTodoByIDService(ctx context.Context, in *pb.TodoQuery) (*pb.Todo, error) {
	todos, err := executeSelectStatement(fmt.Sprintf("SELECT * FROM todo where id = %d", in.Id), s.db)
	if err != nil {
		return nil, err
	}
	return &todos[0], nil
}

func (s *models.server) UpdateTodoService(ctx context.Context, in *pb.Todo) (*pb.Todo, error) {
	err := executeUpdateStatement(in, s.db)
	if err != nil {
		return nil, err
	}
	return in, nil
}
