package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/anpjavatech/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

// to get the database connection.

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("not able to load config file.", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
