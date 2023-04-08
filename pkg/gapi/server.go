package gapi

import (
	"fmt"

	carUsecase "github.com/darwishdev/obd_api/app/car/usecase"
	centerUsecase "github.com/darwishdev/obd_api/app/center/usecase"
	reviewUsecase "github.com/darwishdev/obd_api/app/review/usecase"
	userUsecase "github.com/darwishdev/obd_api/app/user/usecase"
	winchUsecase "github.com/darwishdev/obd_api/app/winch/usecase"
	auth "github.com/darwishdev/obd_api/pkg/auth"
	"github.com/darwishdev/obd_api/pkg/pb/obd/v1/obdv1connect"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"

	"github.com/darwishdev/obd_api/pkg/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	obdv1connect.UnimplementedObdHandler
	config        util.Config
	store         db.Store
	tokenMaker    auth.Maker
	userUsecase   *userUsecase.UserUsecase
	centerUsecase *centerUsecase.CenterUsecase
	reviewUsecase *reviewUsecase.ReviewUsecase
	carUsecase    *carUsecase.CarUsecase
	winchUsecase  *winchUsecase.WinchUsecase
}

// NewServer creates a new gRPC server.
func newServer(config util.Config, store db.Store) (obdv1connect.ObdHandler, error) {
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// appErrors := errors.NewAppErrors()

	server := &Server{
		config:        config,
		store:         store,
		tokenMaker:    tokenMaker,
		userUsecase:   userUsecase.NewUserUsecase(store, tokenMaker, config),
		centerUsecase: centerUsecase.NewCenterUsecase(store),
		carUsecase:    carUsecase.NewCarUsecase(store),
		reviewUsecase: reviewUsecase.NewReviewUsecase(store),
		winchUsecase:  winchUsecase.NewWinchUsecase(store),
	}

	return server, nil
}
