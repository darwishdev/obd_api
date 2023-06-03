package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *SessionRepo) SessionsList(ctx context.Context, req *int64) (*[]db.SessionsListRow, error) {
	sessions, err := repo.store.SessionsList(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
	}
	return &sessions, nil
}

func (repo *SessionRepo) SessionGetCodes(ctx context.Context, req *int64) (*[]db.Code, error) {
	sessions, err := repo.store.SessionGetCodes(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
	}
	return &sessions, nil
}
