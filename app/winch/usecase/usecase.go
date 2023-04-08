package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/winch/factory"
	"github.com/darwishdev/obd_api/app/winch/repo"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type WinchUsecaseInterface interface {
	WinchList(ctx context.Context, req *obdv1.WinchListRequest) (*obdv1.WinchListResponse, error)
}

type WinchUsecase struct {
	repo    repo.WinchRepoInterface
	factory factory.WinchFactoryInterface
}

func NewWinchUsecase(store db.Store) *WinchUsecase {
	repo := repo.NewWinchRepo(store)
	factory := factory.NewWinchFactory()
	return &WinchUsecase{
		repo:    repo,
		factory: factory,
	}
}
