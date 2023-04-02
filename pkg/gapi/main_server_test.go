package gapi

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/darwishdev/obd_api/pkg/pb/obd/v1/obdv1connect"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
	_ "github.com/lib/pq"
)

var testServer obdv1connect.ObdHandler

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../", "test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries := db.NewStore(testDB)

	testServer, err = newServer(config, &testQueries)
	if err != nil {
		log.Fatal("cannot init server:", err)
	}
	os.Exit(m.Run())
}
