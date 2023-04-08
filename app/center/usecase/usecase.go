package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/center/factory"
	"github.com/darwishdev/obd_api/app/center/repo"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CenterUsecaseInterface interface {
	CentersList(ctx context.Context, req *obdv1.CentersListRequest) (*obdv1.CentersListResponse, error)
}

type CenterUsecase struct {
	repo    repo.CenterRepoInterface
	factory factory.CenterFactoryInterface
}

func NewCenterUsecase(store db.Store) *CenterUsecase {
	repo := repo.NewCenterRepo(store)
	factory := factory.NewCenterFactory()
	return &CenterUsecase{
		repo:    repo,
		factory: factory,
	}
}
