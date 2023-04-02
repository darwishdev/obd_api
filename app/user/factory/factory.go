package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type UserFactoryInterface interface {
	NewUserFromProto(req *obdv1.UserCreateRequest) (*db.UserCreateParams, error)
	NewUserCreateFromSqlResponse(req *db.User) (*obdv1.UserCreateResponse, error)
	NewUserLoginFromSqlResponse(req *db.User) (*obdv1.UserLoginResponse, error)
	UserLoginRequestValidation(req *obdv1.UserLoginRequest) error
}

type UserFactory struct {
}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}
