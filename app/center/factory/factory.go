package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CenterFactoryInterface interface {
	NewCentersListFromSqlResponse(req *[]db.Center) (*obdv1.CentersListResponse, error)
}

type CenterFactory struct {
}

func NewCenterFactory() *CenterFactory {
	return &CenterFactory{}
}
