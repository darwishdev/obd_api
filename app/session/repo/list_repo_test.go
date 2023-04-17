package repo

import (
	"context"
	"database/sql"
	"testing"
	"time"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewsList(t *testing.T) {

	// Set up the expected behavior of the mock store
	mockReviews := []db.ReviewsListRow{
		{
			UserName:   "test user",
			CenterName: "test center",
			Review:     "Review 1",
			Rate:       4,
			CreatedAt:  time.Now(),
		},
		{
			UserName:   "test user 2",
			CenterName: "test center 2",
			Review:     "Review 2",
			Rate:       4,
			CreatedAt:  time.Now(),
		},
	}

	req := db.ReviewsListParams{CenterID: sql.NullInt64{Int64: 1}}
	store.EXPECT().ReviewsList(gomock.Any(), req).Return(mockReviews, nil)

	// Create a new instance of the center repository

	// Call the ReviewsList function
	reviews, err := repo.ReviewsList(context.Background(), &req)
	assert.NoError(t, err)
	assert.Equal(t, mockReviews, *reviews)
}
