package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/user/factory"
	"github.com/darwishdev/obd_api/app/user/repo"
	"github.com/darwishdev/obd_api/pkg/auth"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type UserUsecaseInterface interface {
	UserCreate(ctx context.Context, req *obdv1.UserCreateRequest) (*obdv1.UserCreateResponse, error)
	UserAuthorize(ctx context.Context, req int64) (*obdv1.UserAuthorizeResponse, error)
	UserCreateRequestValidation(req *obdv1.UserCreateRequest) (violations []*errdetails.BadRequest_FieldViolation)
	UserUpdate(ctx context.Context, req *obdv1.UserUpdateRequest, userId int64) (*obdv1.UserUpdateResponse, error)
}

type UserUsecase struct {
	repo       repo.UserRepoInterface
	factory    factory.UserFactoryInterface
	tokenMaker auth.Maker
	config     util.Config
}

func NewUserUsecase(store db.Store, tokenMaker auth.Maker, config util.Config) *UserUsecase {
	repo := repo.NewUserRepo(store)
	factory := factory.NewUserFactory()
	return &UserUsecase{
		repo:       repo,
		factory:    factory,
		tokenMaker: tokenMaker,
		config:     config,
	}
}
