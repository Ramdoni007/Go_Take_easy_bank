package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdoni007/Take_Easy_Bank/api"
	db "github.com/ramdoni007/Take_Easy_Bank/db/sqlc"
	"github.com/ramdoni007/Take_Easy_Bank/util"
	"log"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("tidak bisa terhubung ke configLoad ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke database ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke server ", err)
	}
}
