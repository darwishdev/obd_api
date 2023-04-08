package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
var area Area

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../", "test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	area, err = testQueries.AreaCreate(context.Background(), getValidArea())
	if err != nil {
		log.Fatal("cannot create test area:", err)
	}

	os.Exit(m.Run())
}
