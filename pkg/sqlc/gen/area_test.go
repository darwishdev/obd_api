package db

import (
	"context"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type areaCreateTest struct {
	name      string
	area      string
	expectErr bool
}

func getValidArea() string {
	return util.RandomString(10)
}

func TestAreaCreate(t *testing.T) {
	testcases := []areaCreateTest{
		{
			name:      "ValidArea",
			area:      getValidArea(),
			expectErr: false,
		},
		{
			name:      "TooLongName",
			area:      util.RandomString(300),
			expectErr: true,
		},
	}
	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the AreaCreate function with the user data from the current test case
			createdArea, err := testQueries.AreaCreate(context.Background(), tc.area)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}
			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  areas created during test
			testQueries.AreaDelete(context.Background(), createdArea.AreaID)

		})
	}

}
