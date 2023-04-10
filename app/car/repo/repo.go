package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CarRepoInterface interface {
	BrandsList(ctx context.Context) (*[]db.CarBrandsListRow, error)
	CarCreate(ctx context.Context, req *db.CarCreateParams) (*db.Car, error)
	CarUpdate(ctx context.Context, req *db.CarUpdateParams) (*db.Car, error)
}

type CarRepo struct {
	store db.Store
}

func NewCarRepo(store db.Store) *CarRepo {
	return &CarRepo{
		store: store,
	}
}
