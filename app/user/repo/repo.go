package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type UserRepoInterface interface {
	UserCreate(ctx context.Context, req *db.UserCreateParams) (*db.User, error)
	UserGetByUsername(ctx context.Context, req *string) (*db.User, error)
}

type UserRepo struct {
	store db.Store
}

func NewUserRepo(store *db.Store) *UserRepo {
	return &UserRepo{
		store: *store,
	}
}
