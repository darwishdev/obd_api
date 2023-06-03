package factory

import (
	"reflect"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewCenterFromSqlResponse(t *testing.T) {
	c := &db.Center{
		CenterID:  1,
		Name:      "Test Center",
		Phone:     "123456789",
		Address:   "Test Address",
		AreaID:    1,
		Lat:       123.456,
		Long:      456.789,
		CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	expected := &obdv1.Center{
		CenterId:  1,
		Name:      "Test Center",
		Phone:     "123456789",
		Address:   "Test Address",
		AreaId:    1,
		Lat:       123.456,
		Long:      456.789,
		CreatedAt: timestamppb.New(c.CreatedAt),
		DeletedAt: nil,
	}

	result, err := NewCenterFromSqlResponse(c)
	if err != nil {
		t.Fatalf("NewCenterFromSqlResponse error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewCenterFromSqlResponse returned unexpected result: got %+v, expected %+v", result, expected)
	}
}
