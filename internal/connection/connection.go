package connection

import (
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
	query := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", c.User, c.Password, c.Server, c.Port, c.Database)

	db, err := sql.Open("sqlserver", query)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &Server{
		db,
	}, nil
}
