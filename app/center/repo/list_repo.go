package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *CenterRepo) CentersList(ctx context.Context, req *db.CentersListParams) (*[]db.CenterInfo, error) {
	centers, err := repo.store.CentersList(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &centers, nil
}
