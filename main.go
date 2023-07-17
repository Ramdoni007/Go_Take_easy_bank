package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdoni007/Take_Easy_Bank/api"
	db "github.com/ramdoni007/Take_Easy_Bank/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke database ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke server ", err)
	}
}
