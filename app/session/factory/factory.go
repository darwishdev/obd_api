package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type SessionFactoryInterface interface {
	CreateGrpcFromSql(req *db.Session) (*obdv1.SessionCreateResponse, error)
	AttachCodeSqlFromGrpc(req *obdv1.SessionAttachCodeRequest) (*db.SessionAttachCodeParams, error)
	AttachCodeGrpcFromSql(resp *db.SessionCode) (*obdv1.SessionAttachCodeResponse, error)
	AttachValueSqlFromGrpc(req *obdv1.SessionAttachValueRequest) (*db.SessionAttachValueParams, error)
	AttachValueGrpcFromSql(resp *db.SessionValue) (*obdv1.SessionAttachValueResponse, error)
	ListGrpcFromSqlArr(req *[]db.Session) (*obdv1.SessionsListResponse, error)
	CloseGrpcFromSql(req *db.Session) (*obdv1.SessionCloseResponse, error)
}

type SessionFactory struct {
}

func NewSessionFactory() *SessionFactory {
	return &SessionFactory{}
}
