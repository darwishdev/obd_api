package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (repo *UserRepo) UserGetByUsername(ctx context.Context, req *string) (*db.UserInfo, error) {
	product, err := repo.store.UserGetByUsername(context.Background(), *req)
	if err != nil {
		return nil, validator.ParseReadDbErrMsg(err)
	}
	return &product, nil

}
