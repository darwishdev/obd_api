package usecase

import (
	"context"

	"github.com/darwishdev/obd_api/app/review/factory"
	"github.com/darwishdev/obd_api/app/review/repo"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type ReviewUsecaseInterface interface {
	ReviewsList(ctx context.Context, req *obdv1.ReviewsListRequest) (*obdv1.ReviewsListResponse, error)
}

type ReviewUsecase struct {
	repo    repo.ReviewRepoInterface
	factory factory.ReviewFactoryInterface
}

func NewReviewUsecase(store db.Store) *ReviewUsecase {
	repo := repo.NewReviewRepo(store)
	factory := factory.NewReviewFactory()
	return &ReviewUsecase{
		repo:    repo,
		factory: factory,
	}
}
