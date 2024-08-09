package handlers

import (
	"fmt"
	"os"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetEnvVariable(key string) string {
	_ = godotenv.Load()
	return os.Getenv(key)
}

var DBHost = GetEnvVariable("DBHost")
var DBPassword = GetEnvVariable("DBPassword")
var DBUser = GetEnvVariable("DBUser")
var db = newDB()

func newDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=todo sslmode=disable password=%s host=%s", DBUser, DBPassword, DBHost))
	if err != nil {
		return nil
	}
	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil
	}
	return db
}

func executeSelectStatement(sqlStatement string) ([]pb.Todo, error) {
	if db == nil {
		return nil, fmt.Errorf("couldn't connect to the DB at %s:%s", DBHost, "5432")
	}

	returnTodos := []pb.Todo{}
	err := db.Select(&returnTodos, sqlStatement)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return returnTodos, nil
}

func executeInsertStatement(description string) (*pb.Todo, error) {
	if db == nil {
		return nil, fmt.Errorf("couldn't connect to the DB at %s:%s", DBHost, "5432")
	}
	lastInsertId := 0
	err := db.QueryRow("INSERT INTO todo (Description) VALUES($1) RETURNING id", description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int64(lastInsertId), Description: description}, nil
}

func executeDeleteStatement(id int) (*pb.Todo, error) {
	if db == nil {
		return nil, fmt.Errorf("couldn't connect to the DB at %s:%s", DBHost, "5432")
	}
	desc := ""
	err := db.QueryRow("Delete FROM todo WHERE ID=$1 RETURNING description", id).Scan(&desc)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int64(id), Description: desc}, nil
}
func executeUpdateStatement(in *pb.Todo) error {
	if db == nil {
		return fmt.Errorf("couldn't connect to the DB at %s:%s", DBHost, "5432")
	}
	tx := db.MustBegin()
	tx.MustExec("Update todo Set description=$2 WHERE ID=$1", in.Id, in.Description)
	tx.Commit()
	return nil
}
