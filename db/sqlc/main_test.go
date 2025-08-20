package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/MustafaKheda/simplebank/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot Load Config File")
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot Create Connection")
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
