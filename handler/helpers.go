package handler

import (
	"fmt"
	"os"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getEnvVariable(key string) string {
	return os.Getenv(key)
}

func newDB() (*sqlx.DB, error) {
	var DBHost = getEnvVariable("DBHost")
	var DBPassword = getEnvVariable("DBPassword")
	var DBUser = getEnvVariable("DBUser")
	var DBPort = getEnvVariable("DBPort")
	args := fmt.Sprintf("user=%s dbname=todo sslmode=disable password=%s host=%s port=%s", DBUser, DBPassword, DBHost, DBPort)
	db, err := sqlx.Connect("postgres", args)
	if err != nil {
		return nil, err
	}
	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todo (ID Serial PRIMARY KEY, Description TEXT, Titel TEXT)")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func executeSelectStatement(sqlStatement string) ([]pb.Todo, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}

	returnTodos := []pb.Todo{}
	err = db.Select(&returnTodos, sqlStatement)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	if len(returnTodos) == 0 {
		fmt.Println("got 0 results")
	}
	return returnTodos, nil
}

func executeInsertStatement(description string, titel string) (*pb.Todo, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	lastInsertId := 0
	err = db.QueryRow("INSERT INTO todo (Description) VALUES($1) RETURNING id", description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: fmt.Sprint(lastInsertId), Description: description, Titel: titel}, nil
}

func executeDeleteStatement(id string) (*pb.Todo, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	todo := &pb.Todo{}
	err = db.QueryRow("Delete FROM todo WHERE ID=$1 RETURNING *", id).Scan(todo.Id, todo.Description, todo.Titel)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func executeUpdateStatement(in *pb.Todo) error {
	db, err := newDB()
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("Update todo Set description=$2, titel=$3  WHERE ID=$1", in.Id, in.Description, in.Titel)
	tx.Commit()
	return nil
}
