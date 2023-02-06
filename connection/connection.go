package connection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	*sql.DB
}

func init() {
	err := godotenv.Load("local-dev.env")
	if err != nil {
		log.Fatal(err)
	}
}

var (
	database = os.Getenv("DATABASE")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	port     = os.Getenv("PORT")
	server   = os.Getenv("SERVER")
)

func NewConn() (*Server, error) {
	query := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	db, err := sql.Open("sqlserver", query)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	return &Server{
		db,
	}, nil
}
