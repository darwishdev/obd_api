package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type ReviewRepoInterface interface {
	ReviewsList(ctx context.Context, req *db.ReviewsListParams) (*[]db.ReviewsListRow, error)
	ReviewCreate(ctx context.Context, req *db.ReviewCreateParams) (*db.Review, error)
}

type ReviewRepo struct {
	store db.Store
}

func NewReviewRepo(store db.Store) *ReviewRepo {
	return &ReviewRepo{
		store: store,
	}
}
