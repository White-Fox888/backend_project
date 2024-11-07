package conndb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

func InitDB(DB string) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &Database{Conn: conn}, nil
}

func (db *Database) Close() {
	err := db.Conn.Close(context.Background())
	if err != nil {
		fmt.Printf("Error closing connection to database: %v\n", err)
	}
}
