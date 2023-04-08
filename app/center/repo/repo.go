package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CenterRepoInterface interface {
	CentersList(ctx context.Context, req *int64) (*[]db.Center, error)
}

type CenterRepo struct {
	store db.Store
}

func NewCenterRepo(store db.Store) *CenterRepo {
	return &CenterRepo{
		store: store,
	}
}
