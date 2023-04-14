package db

import (
	"context"
	"log"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type centerCreateTest struct {
	name        string
	input       CenterCreateParams
	expectErr   bool
	expectedRes Center
}

func createNewArea() Area {
	area, err := testQueries.AreaCreate(context.Background(), getValidArea())
	if err != nil {
		log.Fatal("cannot create test area:", err)
	}
	return area
}

func createNewCenter() Center {
	var validArea Area = createNewArea()

	validCenter := getValidCenter()

	validCenter.AreaID = validArea.AreaID
	area, err := testQueries.CenterCreate(context.Background(), validCenter)
	if err != nil {
		log.Fatal("cannot create test area:", err)
	}
	return area
}

func getValidCenter() CenterCreateParams {
	address := util.GenerateRandomAddress()
	return CenterCreateParams{
		Name:     util.RandomName(),
		Phone:    util.RandomPhone(),
		Location: util.RandomURL(),
		Address:  util.AddressToString(address),
		Lat:      address.Lat,
		Long:     address.Long,
	}
}
func TestCenterCreate(t *testing.T) {
	var validArea Area = createNewArea()

	validCenter := getValidCenter()
	inValidCenter := getValidCenter()
	overrides := map[string]interface{}{"AreaID": int64(-1)}
	validCenter.AreaID = validArea.AreaID
	err := util.SetStructOverrides(&inValidCenter, overrides)
	if err != nil {
		log.Fatal(err)
	}

	testcases := []centerCreateTest{
		{
			name:      "ValidInput",
			input:     validCenter,
			expectErr: false,
			expectedRes: Center{
				Name:     validCenter.Name,
				Phone:    validCenter.Phone,
				Location: validCenter.Location,
				Address:  validCenter.Address,
				AreaID:   validCenter.AreaID,
				Lat:      validCenter.Lat,
				Long:     validCenter.Long,
			},
		},
		{
			name:        "InvalidAreaID",
			input:       inValidCenter,
			expectErr:   true,
			expectedRes: Center{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			createdCenter, err := testQueries.CenterCreate(context.Background(), tc.input)
			if tc.expectErr && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
			if createdCenter.Name != tc.expectedRes.Name ||
				createdCenter.Phone != tc.expectedRes.Phone ||
				createdCenter.Location != tc.expectedRes.Location ||
				createdCenter.Address != tc.expectedRes.Address ||
				createdCenter.AreaID != tc.expectedRes.AreaID ||
				createdCenter.Lat != tc.expectedRes.Lat ||
				createdCenter.Long != tc.expectedRes.Long {
				t.Errorf("unexpected result: got %v, expected %v", createdCenter, tc.expectedRes)
			}
		})
	}
}

type centersListTest struct {
	name      string
	req       CentersListParams
	expectLen int
}

func TestCentersList(t *testing.T) {
	validCenter := getValidCenter()
	validArea := createNewArea()
	validCenter.AreaID = validArea.AreaID
	// // Create some test centers
	// for i := 1; i <= 5; i++ {
	// 	_, err := testQueries.CenterCreate(context.Background(), validCenter)
	// 	if err != nil {
	// 		t.Fatalf("Failed to create center: %v", err)
	// 	}
	// }

	// Define a slice of test cases
	testcases := []centersListTest{
		{
			name: "ValidAreaID",
			req: CentersListParams{
				Lat:  31.0444,
				Long: 32.2357,
			},
			expectLen: 5,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the CentersList function with the area ID from the current test case
			centers, err := testQueries.CentersList(context.Background(), tc.req)
			if err != nil {
				t.Fatalf("Failed to retrieve centers: %v", err)
			}

			// If the current test case expects a certain number of centers, check that the length of the
			// returned slice matches the expected length
			if len(centers) != tc.expectLen {
				t.Errorf("Expected %d centers but got %d", tc.expectLen, len(centers))
			}
		})
	}
}
