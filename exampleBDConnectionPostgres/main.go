package main

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "test"
)

func main() {
	conf := pgx.ConnConfig{Database: "postgres", Host: "localhost", Port: 5432, User: "test", Password: "test"}
	fmt.Println("open connection")
	conn, err := pgx.Connect(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("ping database")
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("select some rows")
	rows, err := conn.Query("select id, text from test.test")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Println("read that rows")
	for rows.Next() {
		var id int
		var text string
		err = rows.Scan(&id, &text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read rows failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(id, text)
	}

	fmt.Println("search error in rows")
	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "if some error from rows found: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("insert new rows")
	commandTag, err := conn.Exec("INSERT INTO test.test(id, text) VALUES (3, 'new');")
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem with execute: %v\n", err)
		os.Exit(1)
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Fprintf(os.Stderr, "No row found to delete\n")
		os.Exit(1)
	}

	fmt.Println("end it all")
}
