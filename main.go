package main

import (
	"database/sql"
	"log"

	"github.com/anpjavatech/simple_bank/api"
	db "github.com/anpjavatech/simple_bank/db/sqlc"
	"github.com/anpjavatech/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err.Error())
	}
}
