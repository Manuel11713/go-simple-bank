package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Manuel11713/simple-bank/utils"
	_ "github.com/lib/pq"
)

var queries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot conect to db: ", err)
	}

	queries = New(testDB)

	os.Exit(m.Run())
}
