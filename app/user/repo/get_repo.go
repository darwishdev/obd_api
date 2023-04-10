package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (repo *UserRepo) UserGet(ctx context.Context, req *int64) (*db.UserInfo, error) {
	user, err := repo.store.UserGet(context.Background(), *req)
	if err != nil {
		return nil, validator.ParseReadDbErrMsg(err)
	}

	return &user, nil

}
