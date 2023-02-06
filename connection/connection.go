package connection

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/alexparco/FactGo/config"
	_ "github.com/denisenkom/go-mssqldb"
)

type Server struct {
	*sql.DB
}

func NewConn(c *config.Conf) (*Server, error) {
	query := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", c.Server, c.User, c.Password, c.Port, c.Database)

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
