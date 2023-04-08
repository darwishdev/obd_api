package factory

import (
	"fmt"
	"log"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/stretchr/testify/require"
)

func getValidCarCreateReq(field string, value interface{}) *obdv1.CarCreateRequest {
	req := &obdv1.CarCreateRequest{
		CarBrandModelId: 1,
		ModelYear:       int32(util.RandomInt(2010, int64(time.Now().Year()))),
	}

	if field != "" {
		err := util.SetField(req, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}

	return req

}

type carCreateRequestValidationTest struct {
	name        string
	input       *obdv1.CarCreateRequest
	expectError bool
}

func TestCarCreateRequestValidation(t *testing.T) {
	tests := []carCreateRequestValidationTest{
		// Valid user
		{name: "valid", input: getValidCarCreateReq("", ""), expectError: false},
		{name: "invalid model id", input: getValidCarCreateReq("CarBrandModelId", int64(-1)), expectError: true},
		{name: "invalid year too old", input: getValidCarCreateReq("ModelYear", int32(1980)), expectError: true},
		{name: "invalid year after current year", input: getValidCarCreateReq("ModelYear", int32(2030)), expectError: true},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := carCreateRequestValidation(tc.input)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
