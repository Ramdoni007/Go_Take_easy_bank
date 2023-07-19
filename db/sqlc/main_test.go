package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdoni007/Take_Easy_Bank/util"
	"log"
	"os"
	"testing"
)

var testQueris *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Tidak bisa Terhubung ke LoadConfi", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("tidak bisa terhubung ke database ", err)
	}

	testQueris = New(testDB)

	os.Exit(m.Run())
}
