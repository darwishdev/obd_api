package factory

import (
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *UserFactory) UserLoginRequestValidation(req *obdv1.UserLoginRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("email_%w", err))
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("password_%w", err))
	}
	return nil
}

func (f *UserFactory) NewUserLoginFromSqlResponse(resp *db.User) (*obdv1.UserLoginResponse, error) {
	serverResp := &obdv1.UserLoginResponse{
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
