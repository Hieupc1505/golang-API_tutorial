package main

import (
	"context"
	"log"

	"github.com/hieupc09/simple_api/api"
	db "github.com/hieupc09/simple_api/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DB_SOURCE = "postgresql://root:secret@localhost:5432/todo_list?sslmode=disable"
	PORT      = ":3000"
)

func main() {
	coonPool, err := pgxpool.New(context.Background(), DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(coonPool)
	server := api.NewServer(store)

	err = server.Start(PORT)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
