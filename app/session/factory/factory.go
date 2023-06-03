package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type SessionFactoryInterface interface {
	CreateGrpcFromSql(req *db.Session) (*obdv1.SessionCreateResponse, error)
	AttachCodesSqlFromGrpc(req *obdv1.SessionAttachCodesRequest) (*db.SessionAttachCodesParams, error)
	AttachValuesSqlFromGrpc(req *obdv1.SessionAttachValuesRequest) (*db.SessionAttachValuesParams, error)
	ListGrpcFromSqlArr(req *[]db.SessionsListRow) (*obdv1.SessionsListResponse, error)
	CloseGrpcFromSql(req *db.Session) (*obdv1.SessionCloseResponse, error)
	GetCodesGrpcFromSqlArr(req *[]db.Code) (*obdv1.SessionGetCodesResponse, error)
}

type SessionFactory struct {
}

func NewSessionFactory() *SessionFactory {
	return &SessionFactory{}
}
