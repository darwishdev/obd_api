package factory

import (
	"fmt"
	"log"
	"testing"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/stretchr/testify/require"
)

func getValidReviewCreateReq(field string, value interface{}) *obdv1.ReviewCreateRequest {
	req := &obdv1.ReviewCreateRequest{
		CenterId: 1,
		Review:   util.RandomString(10),
		Rate:     4,
	}

	if field != "" {
		err := util.SetField(req, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}

	return req

}

type reviewCreateRequestValidationTest struct {
	name        string
	input       *obdv1.ReviewCreateRequest
	expectError bool
}

func TestReviewCreateRequestValidation(t *testing.T) {
	tests := []reviewCreateRequestValidationTest{
		// Valid user
		{name: "valid", input: getValidReviewCreateReq("", ""), expectError: false},
		{name: "too short review", input: getValidReviewCreateReq("Review", "a"), expectError: true},
		{name: "too long review", input: getValidReviewCreateReq("Review", util.RandomString(300)), expectError: true},
		{name: "invalid rate too small", input: getValidReviewCreateReq("Rate", int64(0)), expectError: true},
		{name: "invalid rate too long", input: getValidReviewCreateReq("Rate", int64(6)), expectError: true},
		{name: "invalid center", input: getValidReviewCreateReq("CenterId", int64(-1)), expectError: true},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := reviewCreateRequestValidation(tc.input)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
