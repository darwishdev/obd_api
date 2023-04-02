package gapi

import (
	"fmt"

	userUsecase "github.com/darwishdev/obd_api/app/user/usecase"
	auth "github.com/darwishdev/obd_api/pkg/auth"
	"github.com/darwishdev/obd_api/pkg/pb/obd/v1/obdv1connect"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"

	"github.com/darwishdev/obd_api/pkg/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	obdv1connect.UnimplementedObdHandler
	config      util.Config
	store       *db.Store
	tokenMaker  auth.Maker
	userUsecase *userUsecase.UserUsecase
}

// NewServer creates a new gRPC server.
func newServer(config util.Config, store *db.Store) (obdv1connect.ObdHandler, error) {
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// appErrors := errors.NewAppErrors()

	server := &Server{
		config:      config,
		store:       store,
		tokenMaker:  tokenMaker,
		userUsecase: userUsecase.NewUserUsecase(store, tokenMaker, config),
	}

	return server, nil
}
