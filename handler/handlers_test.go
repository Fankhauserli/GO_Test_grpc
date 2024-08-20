package handler

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"google.golang.org/grpc"
)

func TestServer_CreateTodoService(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.TodoRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.TodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateTodoService(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateTodoService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteTodoService(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.TodoQuery
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteTodoService(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteTodoService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetAllTodosService(t *testing.T) {
	type args struct {
		in     *pb.Null
		stream grpc.ServerStreamingServer[pb.Todo]
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GetAllTodosService(tt.args.in, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetAllTodosService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetTodoByIDService(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.TodoQuery
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTodoByIDService(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetTodoByIDService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetTodoByIDService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateTodoService(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Todo
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateTodoService(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateTodoService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}
