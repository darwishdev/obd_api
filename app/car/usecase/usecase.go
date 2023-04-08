package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/car/factory"
	"github.com/darwishdev/obd_api/app/car/repo"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CarUsecaseInterface interface {
	CarCreate(ctx context.Context, req *obdv1.CarCreateRequest, userId int64) (*obdv1.CarCreateResponse, error)
}

type CarUsecase struct {
	repo    repo.CarRepoInterface
	factory factory.CarFactoryInterface
}

func NewCarUsecase(store db.Store) *CarUsecase {
	repo := repo.NewCarRepo(store)
	factory := factory.NewCarFactory()
	return &CarUsecase{
		repo:    repo,
		factory: factory,
	}
}
