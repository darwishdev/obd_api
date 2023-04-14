package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *SessionRepo) SessionCreate(ctx context.Context, req *db.SessionCreateParams) (*db.Session, error) {
	resp, err := repo.store.SessionCreate(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &resp, nil
}

func (repo *SessionRepo) SessionAttachCode(ctx context.Context, req *db.SessionAttachCodeParams) (*db.SessionCode, error) {
	resp, err := repo.store.SessionAttachCode(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &resp, nil
}

func (repo *SessionRepo) SessionAttachValue(ctx context.Context, req *db.SessionAttachValueParams) (*db.SessionValue, error) {
	resp, err := repo.store.SessionAttachValue(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &resp, nil
}
