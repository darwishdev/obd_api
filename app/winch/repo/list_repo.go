package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *WinchRepo) WinchList(ctx context.Context, req *int64) (*[]db.WinchListRow, error) {
	winches, err := repo.store.WinchList(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &winches, nil
}
