package factory

import (
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userUpdateRequestValidation(req *obdv1.UserUpdateRequest) error {
	if req.Name != "" {
		if err := validator.ValidateString(req.Name, 3, 200); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("name_%w", err))
		}
	}
	if req.Email != "" {
		if err := validator.ValidateEmail(req.Email); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("email_%w", err))
		}
	}
	if req.Phone != "" {
		if err := validator.ValidatePhone(req.Phone); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("phone_%w", err))
		}
	}
	if req.Password != "" {
		if err := validator.ValidatePassword(req.Password); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("password_%w", err))
		}
	}
	return nil
}

func (f *UserFactory) NewUserUpdateFromProto(req *obdv1.UserUpdateRequest) (*db.UserUpdateParams, error) {
	if err := userUpdateRequestValidation(req); err != nil {
		return nil, err
	}
	var hashedPassword string
	var err error
	if req.Password != "" {
		hashedPassword, err = util.HashPassword(req.Password)
		if err != nil {
			return nil, validator.InternalErr(fmt.Errorf("failed to hash password %w", err))
		}
	}

	resp := &db.UserUpdateParams{
		Name:     util.ToNullString(req.Name),
		Phone:    util.ToNullString(req.Phone),
		Email:    util.ToNullString(req.Email),
		Password: util.ToNullString(hashedPassword),
	}
	return resp, nil
}

func (f *UserFactory) NewUserUpdateFromSqlResponse(resp *db.User) (*obdv1.UserUpdateResponse, error) {
	serverResp := &obdv1.UserUpdateResponse{
		User: &obdv1.User{
			Name:      resp.Name,
			Phone:     resp.Phone,
			Email:     resp.Email,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			DeletedAt: timestamppb.New(resp.DeletedAt.Time),
		},
	}
	return serverResp, nil
}
