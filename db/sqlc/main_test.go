package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable"
)

var testQueris *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke database ", err)
	}

	testQueris = New(conn)

	os.Exit(m.Run())
}
