package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type WinchFactoryInterface interface {
	ListGrpcFromSqlArr(req *[]db.WinchInfo) (*obdv1.WinchListResponse, error)
	ListSqlFromGrpc(req *obdv1.WinchListRequest) (*db.WinchListParams, error)
}

type WinchFactory struct {
}

func NewWinchFactory() *WinchFactory {
	return &WinchFactory{}
}
