package main

import (
	"context"
	"database/sql"
	"dayone/gensql/postgresdb"
	"fmt"

	// PostgreSQL driver
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func main() {

	// getUser := adapter.NewAdapter()

	// user := user.NewUser(getUser)
	// fmt.Println(user.GetUsername("1"))
	// fmt.Println(user.GetUsername("999"))

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=example dbname=testdb sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Database connection established successfully")
	ctx := context.Background()
	pgxConnection, err := pgx.ConnectWithOptions(ctx, "postgres://postgres:example@localhost:5432/testdb?sslmode=disable",
		pgx.ParseConfigOptions{},
	)
	if err != nil {
		fmt.Println("Error connecting to the database with pgx:", err)
		return
	}
	defer pgxConnection.Close(ctx)
	pdb := postgresdb.New(pgxConnection)

	data, err := pdb.CreateAuthor(ctx, postgresdb.CreateAuthorParams{})
	if err != nil {
		fmt.Println("Error creating author:", err)
		return
	}
	fmt.Println("Author created successfully:", data)

	// postgresdb.New(postgresdb.New(db))

}
