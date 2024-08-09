package models

import (
	pb "github.com/Fankhauserli/GO_Test_grpc/services/models"
	"github.com/jmoiron/sqlx"
)

type server struct {
	pb.UnimplementedToDoServer
	db *sqlx.DB
}
