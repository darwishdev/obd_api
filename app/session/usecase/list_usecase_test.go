package usecase

import (
	"context"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	mockdb "github.com/darwishdev/obd_api/pkg/sqlc/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)
	usecase := NewReviewUsecase(store)

	// Set up the expected behavior of the mock store
	mockReviews := []db.ReviewsListRow{
		{
			UserName:   "test user",
			CenterName: "test review",
			Review:     "Review 1",
			Rate:       4,
			CreatedAt:  time.Now(),
		},
		{
			UserName:   "test user 2",
			CenterName: "test review 2",
			Review:     "Review 2",
			Rate:       4,
			CreatedAt:  time.Now(),
		},
	}
	req := &obdv1.ReviewsListRequest{CenterId: 1}

	store.EXPECT().ReviewsList(gomock.Any(), gomock.Any()).Return(mockReviews, nil)
	reviews, err := usecase.ReviewsList(context.Background(), req)
	assert.NoError(t, err)

	assert.NotNil(t, reviews)
	assert.Equal(t, len(mockReviews), len(reviews.Reviews))
}
