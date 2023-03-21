package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnection() {
	databaseUrl := "postgres://postgres:1234@localhost:5432/Personal-Web"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {

		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Success connect to database")
}
