package factory

import (
	"reflect"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestListGrpcFromSql(t *testing.T) {
	c := &db.WinchListRow{
		WinchID:     1,
		Name:        "Test Winch",
		Phone:       "123456789",
		DriverPhone: "Test Address",
		CreatedAt:   time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	expected := &obdv1.Winch{
		WinchId:     1,
		Name:        "Test Winch",
		Phone:       "123456789",
		DriverPhone: "Test Address",
		CreatedAt:   timestamppb.New(c.CreatedAt),
	}

	result, err := listGrpcFromSql(c)
	if err != nil {
		t.Fatalf("listGrpcFromSql error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("listGrpcFromSql returned unexpected result: got %+v, expected %+v", result, expected)
	}
}
