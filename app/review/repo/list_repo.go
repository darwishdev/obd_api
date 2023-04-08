package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/lib/pq"
)

func (repo *ReviewRepo) ReviewsList(ctx context.Context, req *db.ReviewsListParams) (*[]db.ReviewsListRow, error) {
	reviews, err := repo.store.ReviewsList(context.Background(), *req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, validator.ParseWriteDbErrMsg(pqErr)
		}
		return nil, validator.ParseReadDbErrMsg(err)
		// return nil, validator.InternalErr(fmt.Errorf("internal_error_%w", err))
	}
	return &reviews, nil
}
