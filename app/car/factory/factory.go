package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type CarFactoryInterface interface {
	CreateSqlFromGrpc(req *obdv1.CarCreateRequest) (*db.CarCreateParams, error)
	CreateGrpcFromSql(req *db.Car) (*obdv1.CarCreateResponse, error)
	UpdateSqlFromGrpc(req *obdv1.CarUpdateRequest) (*db.CarUpdateParams, error)
	UpdateGrpcFromSql(req *db.Car) (*obdv1.CarUpdateResponse, error)
	BrandsListGrpcFromSqlArr(req *[]db.CarBrandsListRow) (*obdv1.BrandsListResponse, error)
}

type CarFactory struct {
}

func NewCarFactory() *CarFactory {
	return &CarFactory{}
}
