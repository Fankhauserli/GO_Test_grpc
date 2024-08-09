package handler

import (
	"fmt"
	"os"

	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnvVariable(key string) string {
	_ = godotenv.Load()
	return os.Getenv(key)
}

func newDB() (*sqlx.DB, error) {
	var DBHost = getEnvVariable("DBHost")
	var DBPassword = getEnvVariable("DBPassword")
	var DBUser = getEnvVariable("DBUser")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=todo sslmode=disable password=%s host=%s", DBUser, DBPassword, DBHost))
	if err != nil {
		return nil, err
	}
	// Test the connection to the database
	if err := db.Ping(); err != nil {
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

func executeInsertStatement(description string) (*pb.Todo, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	lastInsertId := 0
	err = db.QueryRow("INSERT INTO todo (Description) VALUES($1) RETURNING id", description).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int32(lastInsertId), Description: description}, nil
}

func executeDeleteStatement(id int) (*pb.Todo, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	desc := ""
	err = db.QueryRow("Delete FROM todo WHERE ID=$1 RETURNING description", id).Scan(&desc)
	if err != nil {
		return nil, err
	}
	return &pb.Todo{Id: int32(id), Description: desc}, nil
}
func executeUpdateStatement(in *pb.Todo) error {
	db, err := newDB()
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("Update todo Set description=$2 WHERE ID=$1", in.Id, in.Description)
	tx.Commit()
	return nil
}
