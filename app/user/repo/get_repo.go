package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
)

func (repo *UserRepo) UserGetByUsername(ctx context.Context, req *string) (*db.User, error) {
	product, err := repo.store.UserGetByUsername(context.Background(), *req)
	log.Debug().Str("err", "err").Msg("remove me")
	if err != nil {
		return nil, validator.ParseReadDbErrMsg(err)
	}
	return &product, nil

}
