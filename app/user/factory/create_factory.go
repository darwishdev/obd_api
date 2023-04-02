package factory

import (
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userCreateRequestValidation(req *obdv1.UserCreateRequest) error {
	if err := validator.ValidateString(req.Name, 3, 200); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("name_%w", err))
	}
	if err := validator.ValidateEmail(req.Email); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("email_%w", err))
	}
	if err := validator.ValidatePhone(req.Phone); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("phone_%w", err))
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("password_%w", err))
	}
	return nil
}

func (f UserFactory) NewUserFromProto(req *obdv1.UserCreateRequest) (*db.UserCreateParams, error) {
	if err := userCreateRequestValidation(req); err != nil {
		return nil, err
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, validator.InternalErr(fmt.Errorf("failed to hash password %w", err))
	}

	resp := &db.UserCreateParams{
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return resp, nil
}

func (f *UserFactory) NewUserCreateFromSqlResponse(resp *db.User) (*obdv1.UserCreateResponse, error) {
	serverResp := &obdv1.UserCreateResponse{
		User: &obdv1.User{
			Name:      resp.Name,
			Phone:     resp.Phone,
			Email:     resp.Email,
			Password:  resp.Password,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			DeleteAt:  timestamppb.New(resp.DeletedAt.Time),
		},
	}
	return serverResp, nil
}
