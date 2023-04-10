package db

import (
	"context"
	"log"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/stretchr/testify/assert"
)

type carCreateTest struct {
	name      string
	input     CarCreateParams
	expectErr bool
}

// func createNewCar() Car {

// 	validCar := getValidCar()

// 	car, err := testQueries.CarCreate(context.Background(), validCar)
// 	if err != nil {
// 		log.Fatal("cannot create test area:", err)
// 	}
// 	return car
// }

func getValidCar() CarCreateParams {
	user := createNewUser()
	return CarCreateParams{
		CarBrandModelID: 1,
		UserID:          user.UserID,
		ModelYear:       2010,
	}
}
func TestCarCreate(t *testing.T) {
	validCar := getValidCar()
	inValidCar := getValidCar()
	overrides := map[string]interface{}{"CarBrandModelID": int64(-1)}
	err := util.SetStructOverrides(&inValidCar, overrides)
	if err != nil {
		log.Fatal(err)
	}

	testcases := []carCreateTest{
		{
			name:      "ValidInput",
			input:     validCar,
			expectErr: false,
		},
		{
			name:      "InvalidAreaID",
			input:     inValidCar,
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			createdCar, err := testQueries.CarCreate(context.Background(), tc.input)
			if tc.expectErr {
				assert.Error(t, err)
				assert.Empty(t, createdCar)
			}
			if !tc.expectErr {
				assert.NotEmpty(t, createdCar)
				assert.NoError(t, err)

			}
		})
	}
}
