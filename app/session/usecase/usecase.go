package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/session/factory"
	"github.com/darwishdev/obd_api/app/session/repo"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type SessionUsecaseInterface interface {
	SessionCreate(ctx context.Context, req *obdv1.SessionCreateRequest, userId int64) (*obdv1.SessionCreateResponse, error)
	SessionAttachCodes(ctx context.Context, req *obdv1.SessionAttachCodesRequest) (*obdv1.SessionAttachCodesResponse, error)
	SessionAttachValues(ctx context.Context, req *obdv1.SessionAttachValuesRequest) (*obdv1.SessionAttachValuesResponse, error)
	SessionsList(ctx context.Context, req *obdv1.SessionsListRequest) (*obdv1.SessionsListResponse, error)
	SessionGetCodes(ctx context.Context, carId *int64) (*obdv1.SessionGetCodesResponse, error)
}

type SessionUsecase struct {
	repo    repo.SessionRepoInterface
	factory factory.SessionFactoryInterface
}

func NewSessionUsecase(store db.Store) *SessionUsecase {
	repo := repo.NewSessionRepo(store)
	factory := factory.NewSessionFactory()
	return &SessionUsecase{
		repo:    repo,
		factory: factory,
	}
}
