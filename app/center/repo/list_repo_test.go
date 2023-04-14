package repo

// func TestCentersList(t *testing.T) {

// 	// Set up the expected behavior of the mock store
// 	mockCenters := []db.Center{
// 		{
// 			CenterID:  1,
// 			Name:      "Center 1",
// 			Phone:     "123456",
// 			Location:  "location 1",
// 			Address:   "address 1",
// 			AreaID:    1,
// 			Lat:       123.456,
// 			Long:      456.789,
// 			CreatedAt: time.Now(),
// 		},
// 		{
// 			CenterID:  2,
// 			Name:      "Center 2",
// 			Phone:     "789012",
// 			Location:  "location 2",
// 			Address:   "address 2",
// 			AreaID:    1,
// 			Lat:       123.456,
// 			Long:      456.789,
// 			CreatedAt: time.Now(),
// 		},
// 	}
// 	store.EXPECT().CentersList(gomock.Any(), int64(1)).Return(mockCenters, nil)

// 	// Create a new instance of the center repository

// 	// Call the CentersList function
// 	req := int64(1)
// 	centers, err := repo.CentersList(context.Background(), &req)
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockCenters, *centers)
// }
