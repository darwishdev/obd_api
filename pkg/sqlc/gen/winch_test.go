package db

import (
	"context"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type wichsListTest struct {
	name      string
	areaID    int64
	expectLen int
}

func getValidWinchCreate() WinchCreateParams {
	return WinchCreateParams{
		Name:        util.RandomName(),
		Phone:       util.RandomPhone(),
		DriverName:  util.RandomName(),
		DriverPhone: util.RandomPhone(),
	}
}
func TestWinchList(t *testing.T) {
	validWinch := getValidWinchCreate()
	validArea := createNewArea()
	validWinch.AreaID = validArea.AreaID
	// Create some test wichs
	for i := 1; i <= 5; i++ {
		_, err := testQueries.WinchCreate(context.Background(), validWinch)
		if err != nil {
			t.Fatalf("Failed to create wich: %v", err)
		}
	}

	// Define a slice of test cases
	testcases := []wichsListTest{
		{
			name:      "ValidAreaID",
			areaID:    validArea.AreaID,
			expectLen: 5,
		},
		{
			name:      "NonExistentAreaID",
			areaID:    area.AreaID + 10,
			expectLen: 0,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the WinchList function with the area ID from the current test case
			wichs, err := testQueries.WinchList(context.Background(), tc.areaID)
			if err != nil {
				t.Fatalf("Failed to retrieve wichs: %v", err)
			}

			// If the current test case expects a certain number of wichs, check that the length of the
			// returned slice matches the expected length
			if len(wichs) != tc.expectLen {
				t.Errorf("Expected %d wichs but got %d", tc.expectLen, len(wichs))
			}
		})
	}
}
