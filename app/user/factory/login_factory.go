package factory

import (
	"fmt"

	obdv1Car "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	obdv1User "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *UserFactory) UserLoginRequestValidation(req *obdv1User.UserLoginRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("email_%w", err))
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("password_%w", err))
	}
	return nil
}

func (f *UserFactory) LoginGrpcFromSql(resp *db.UserInfo) (*obdv1User.UserLoginResponse, error) {
	serverResp := &obdv1User.UserLoginResponse{
		User: &obdv1User.User{
			Name:      resp.Name,
			Phone:     resp.Phone,
			Email:     resp.Email,
			Password:  resp.Password,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			DeletedAt: timestamppb.New(resp.DeletedAt.Time),
		},
	}
	if resp.BrandName.Valid {
		carInfo := &obdv1Car.CarView{
			BrandName:      resp.BrandName.String,
			BrandModelName: resp.BrandModelName.String,
			ModelYear:      resp.ModelYear.Int32,
			CreatedAt:      timestamppb.New(resp.CarCreatedAt.Time),
		}
		serverResp.Car = carInfo
	}
	return serverResp, nil
}
