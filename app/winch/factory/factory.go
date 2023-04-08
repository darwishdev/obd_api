package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type WinchFactoryInterface interface {
	ListGrpcFromSqlArr(req *[]db.WinchListRow) (*obdv1.WinchListResponse, error)
}

type WinchFactory struct {
}

func NewWinchFactory() *WinchFactory {
	return &WinchFactory{}
}
