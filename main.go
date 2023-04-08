package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"database/sql"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	gapi "github.com/darwishdev/obd_api/pkg/gapi"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	loggerAdapter := zerologadapter.New(log.Logger)
	conn = sqldblogger.OpenDriver(config.DBSource, conn.Driver(), loggerAdapter)
	store := db.NewStore(conn)
	runGrpcHttpServer(config, store)

}
func runGrpcHttpServer(config util.Config, store db.Store) {
	httpServer, err := gapi.NewGrpcHttpServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initialize grpc ,HTTP api server:")
	}
	log.Debug().Str("running on %s", config.GRPCServerAddress).Msg("successfully running")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP listen and serve")
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("HTTP shutdown")
	}
}
