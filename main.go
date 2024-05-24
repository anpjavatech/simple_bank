package main

import (
	"database/sql"
	"log"

	"github.com/anpjavatech/simple_bank/api"
	db "github.com/anpjavatech/simple_bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://anoop:root@localhost:5433/simple-bank?sslmode=disable"
	address  = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server:", err.Error())
	}
}
