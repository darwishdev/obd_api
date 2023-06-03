package usecase

import (
	"context"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCentersList(t *testing.T) {

	// Set up the expected behavior of the mock store
	mockCenters := []db.Center{
		{
			CenterID:  1,
			Name:      "Center 1",
			Phone:     "123456",
			Address:   "address 1",
			AreaID:    1,
			Lat:       123.456,
			Long:      456.789,
			CreatedAt: time.Now(),
		},
		{
			CenterID:  2,
			Name:      "Center 2",
			Phone:     "789012",
			Address:   "address 2",
			AreaID:    1,
			Lat:       123.456,
			Long:      456.789,
			CreatedAt: time.Now(),
		},
	}
	store.EXPECT().CentersList(gomock.Any(), int64(1)).Return(mockCenters, nil)

	// Create a new instance of the center repository

	// Call the CentersList function
	req := &obdv1.CentersListRequest{AreaId: 1}
	centers, err := usecase.CentersList(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, centers)
	assert.Equal(t, len(mockCenters), len(centers.Centers))
	// assert.Equal(t, mockCenters, (*centers).Centers)
}
