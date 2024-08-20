package handler

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setEnvVar(key string, value string) error {
	return os.Setenv(key, value)
}

func getPGContainer(ctx context.Context) (*postgres.PostgresContainer, error) {

	return postgres.Run(ctx,
		"postgres:latest",
		postgres.WithInitScripts(filepath.Join("..", "DB", "tests.sql")),
		postgres.WithDatabase("todo"),
		postgres.WithUsername("gopher"),
		postgres.WithPassword("golang1234"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
}

func setPGEnvVar(ctx context.Context, pgContainer *postgres.PostgresContainer) error {
	host, err := pgContainer.Host(ctx)
	if err != nil {
		return err
	}
	end, err := pgContainer.Endpoint(ctx, "")
	if err != nil {
		return err
	}
	port := strings.Split(end, ":")[1]

	envVar := []struct {
		key   string
		value string
	}{
		{
			key:   "DBHost",
			value: host,
		},
		{
			key:   "DBPassword",
			value: "golang1234",
		},
		{
			key:   "DBUser",
			value: "gopher",
		},
		{
			key:   "DBPort",
			value: port,
		},
	}
	for _, env := range envVar {
		err := setEnvVar(env.key, env.value)
		if err != nil {
			return err
		}
	}
	return nil
}

func Test_executeSelectStatement(t *testing.T) {

	ctx := context.Background()

	pgContainer, err := getPGContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	type args struct {
		sqlStatement string
	}
	tests := []struct {
		name    string
		args    args
		want    []pb.Todo
		wantErr bool
	}{
		{
			name: "Select *",
			args: args{sqlStatement: "SELECT * FROM todo"},
			want: []pb.Todo{
				{Id: "1", Description: "test", Titel: "Test1"},
				{Id: "2", Description: "test2", Titel: "Test2"},
			},
			wantErr: false,
		},
		{
			name: "Select with ID 1",
			args: args{sqlStatement: "SELECT * FROM todo where id = 1"},
			want: []pb.Todo{
				{Id: "1", Description: "test", Titel: "Test1"},
			},
			wantErr: false,
		},
		{
			name:    "Select with wrong SQL",
			args:    args{sqlStatement: "SELECT * FROM todos"},
			want:    nil,
			wantErr: true,
		},
	}

	err = setPGEnvVar(ctx, pgContainer)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeSelectStatement(tt.args.sqlStatement)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeSelectStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeSelectStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeInsertStatement(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := getPGContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	type args struct {
		description string
		titel       string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Todo
		wantErr bool
	}{
		{
			name: "test insert",
			args: args{titel: "Test", description: "test"},
			want: &pb.Todo{
				Id:          "1",
				Titel:       "Test",
				Description: "test",
			},
			wantErr: false,
		},
	}

	err = setPGEnvVar(ctx, pgContainer)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeInsertStatement(tt.args.description, tt.args.titel)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeInsertStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeInsertStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeDeleteStatement(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := getPGContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	err = setPGEnvVar(ctx, pgContainer)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeDeleteStatement(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeDeleteStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeDeleteStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeUpdateStatement(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := getPGContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	type args struct {
		in *pb.Todo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	err = setPGEnvVar(ctx, pgContainer)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := executeUpdateStatement(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("executeUpdateStatement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
