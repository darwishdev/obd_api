package factory

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestListGrpcFromSql(t *testing.T) {
	c := &db.ReviewsListRow{
		UserName:   "Test user",
		CenterName: "Test Center",
		Review:     "test review",
		Rate:       int16(4),
		CreatedAt:  time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	expected := &obdv1.ReviewListRow{
		UserName:   "Test user",
		CenterName: "Test Center",
		Review:     "test review",
		Rate:       int32(4),
		CreatedAt:  timestamppb.New(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
	}

	result, err := listGrpcFromSql(c)
	if err != nil {
		t.Fatalf("listGrpcFromSql error: %v", err)
	}
	if result.UserName != expected.UserName ||
		result.Review != expected.Review ||
		result.Rate != expected.Rate ||
		result.CenterName != expected.CenterName {
		t.Errorf("unexpected result: got %v, expected %v", result, expected)
	}
}

func TestNewReviewsListFromProto(t *testing.T) {
	c := &obdv1.ReviewsListRequest{
		CenterId: 1,
	}

	expected :=
		&db.ReviewsListParams{
			CenterID: sql.NullInt64{Int64: 1, Valid: true},
		}

	result, err := factory.ListSqlFromGrpc(c)
	if err != nil {
		t.Fatalf("listGrpcFromSql error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("listGrpcFromSql returned unexpected result: got %+v, expected %+v", result, expected)
	}
}

func TestListGrpcFromSqlArr(t *testing.T) {
	c := &[]db.ReviewsListRow{
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review",
			Rate:       int16(4),
			CreatedAt:  time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review2",
			Rate:       int16(4),
			CreatedAt:  time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review3",
			Rate:       int16(4),
			CreatedAt:  time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	expected := []obdv1.ReviewListRow{
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review",
			Rate:       int32(4),
			CreatedAt:  timestamppb.New(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
		},
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review2",
			Rate:       int32(4),
			CreatedAt:  timestamppb.New(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
		},
		{

			UserName:   "Test user",
			CenterName: "Center Name",
			Review:     "test review3",
			Rate:       int32(4),
			CreatedAt:  timestamppb.New(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
		},
	}

	result, err := factory.ListGrpcFromSqlArr(c)
	if err != nil {
		t.Fatalf("listGrpcFromSql error: %v", err)
	}

	for i := range expected {
		assert.Equal(t, expected[i].CenterName, result.Reviews[i].CenterName)
		assert.Equal(t, expected[i].UserName, result.Reviews[i].UserName)
		assert.Equal(t, expected[i].Review, result.Reviews[i].Review)
		assert.Equal(t, expected[i].Rate, result.Reviews[i].Rate)
		assert.Equal(t, expected[i].CreatedAt, result.Reviews[i].CreatedAt)
	}

}

type reviewsListRequestValidationTest struct {
	name        string
	input       int64
	expectError bool
}

func TestReviewsListRequestValidation(t *testing.T) {
	tests := []reviewsListRequestValidationTest{
		// Valid user
		{name: "valid", input: 1, expectError: false},
		{name: "invalid name", input: -1, expectError: true},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := reviewsListRequestValidation(&obdv1.ReviewsListRequest{CenterId: tc.input})
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
