package db

import (
	"context"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type reviewCreateTest struct {
	name        string
	input       ReviewCreateParams
	expectErr   bool
	expectedRes Review
}

func getValidReview() ReviewCreateParams {
	validUser := createNewUser()
	validCenter := createNewCenter()
	return ReviewCreateParams{
		CenterID: validCenter.CenterID,
		UserID:   validUser.UserID,
		Review:   util.RandomString(10),
		Rate:     int16(util.RandomInt(1, 5)),
	}
}

func getInValidReview() ReviewCreateParams {
	validUser := createNewUser()
	return ReviewCreateParams{
		CenterID: -1,
		UserID:   validUser.UserID,
		Review:   util.RandomString(10),
		Rate:     int16(util.RandomInt(1, 5)),
	}
}
func TestReviewsCreate(t *testing.T) {
	validReview := getValidReview()

	testcases := []reviewCreateTest{
		{
			name:      "ValidInput",
			input:     validReview,
			expectErr: false,
			expectedRes: Review{
				CenterID: validReview.CenterID,
				UserID:   validReview.UserID,
				Review:   validReview.Review,
				Rate:     validReview.Rate,
			},
		},
		{
			name:        "InvalidCenterID",
			input:       getInValidReview(),
			expectErr:   true,
			expectedRes: Review{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			createdReview, err := testQueries.ReviewCreate(context.Background(), tc.input)
			if tc.expectErr && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
			if createdReview.UserID != tc.expectedRes.UserID ||
				createdReview.Review != tc.expectedRes.Review ||
				createdReview.Rate != tc.expectedRes.Rate ||
				createdReview.CenterID != tc.expectedRes.CenterID {
				t.Errorf("unexpected result: got %v, expected %v", createdReview, tc.expectedRes)
			}
		})
	}
}
