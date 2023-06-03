package db

import (
	"context"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type wichsListTest struct {
	name      string
	req       WinchListParams
	expectLen int
}

func getValidWinchCreate() WinchCreateParams {
	return WinchCreateParams{
		Name:        util.RandomName(),
		Phone:       util.RandomPhone(),
		DriverName:  util.RandomName(),
		DriverPhone: util.RandomPhone(),
		Lat:         util.RandomLatLong(),
		Long:        util.RandomLatLong(),
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
			name: "ValidAreaID",
			req: WinchListParams{
				InLat:  float64(util.RandomLatLong()),
				InLong: float64(util.RandomLatLong()),
			},
			expectLen: 5,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the WinchList function with the area ID from the current test case
			winchs, err := testQueries.WinchList(context.Background(), tc.req)
			if err != nil {
				t.Fatalf("Failed to retrieve winchs: %v", err)
			}

			// If the current test case expects a certain number of winchs, check that the length of the
			// returned slice matches the expected length
			if len(winchs) != tc.expectLen {
				t.Errorf("Expected %d winchs but got %d", tc.expectLen, len(winchs))
			}
		})
	}

	// clean up
	testQueries.WinchClean(context.Background())

}
