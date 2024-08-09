package main

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

func newDB() *sqlx.DB {
	var DBHost = GetEnvVariable("DBHost")
	var DBPassword = GetEnvVariable("DBPassword")
	var DBUser = GetEnvVariable("DBUser")
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

func executeSelectStatement(sqlStatement string, db *sqlx.DB) ([]pb.Todo, error) {

	returnTodos := []pb.Todo{}
	err := db.Select(&returnTodos, sqlStatement)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return returnTodos, nil
}

func executeInsertStatement(description string, db *sqlx.DB) (*pb.Todo, error) {
	lastInsertId := 0
	err := db.QueryRow("INSERT INTO todo (Description) VALUES($1) RETURNING id", description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int64(lastInsertId), Description: description}, nil
}

func executeDeleteStatement(id int, db *sqlx.DB) (*pb.Todo, error) {
	desc := ""
	err := db.QueryRow("Delete FROM todo WHERE ID=$1 RETURNING description", id).Scan(&desc)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int64(id), Description: desc}, nil
}
func executeUpdateStatement(in *pb.Todo, db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec("Update todo Set description=$2 WHERE ID=$1", in.Id, in.Description)
	tx.Commit()
	return nil
}
