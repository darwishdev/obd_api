package repo

import (
	"context"
	"fmt"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *UserRepo) UserCreate(ctx context.Context, req *db.UserCreateParams) (*db.UserInfo, error) {
	user, err := repo.store.UserCreate(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}

	return &user, nil
}
