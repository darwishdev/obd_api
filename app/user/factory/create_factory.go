package factory

import (
	"fmt"

	obdv1car "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
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

func (f *UserFactory) NewUserFromProto(req *obdv1.UserCreateRequest) (*db.UserCreateParams, error) {
	if err := userCreateRequestValidation(req); err != nil {
		return nil, err
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, validator.InternalErr(fmt.Errorf("failed to hash password %w", err))
	}

	resp := &db.UserCreateParams{
		NameArg:            req.Name,
		PhoneArg:           req.Phone,
		EmailArg:           req.Email,
		PasswordArg:        hashedPassword,
		CarBrandModelIDArg: int32(req.CarBrandModel),
		ModelYearArg:       int32(req.CarYear),
	}
	return resp, nil
}

func (f *UserFactory) NewUserCreateFromSqlResponse(resp *db.UserInfo) (*obdv1.UserCreateResponse, error) {
	serverResp := &obdv1.UserCreateResponse{
		User: &obdv1.User{
			UserId:    int32(resp.UserID),
			Name:      resp.Name,
			Phone:     resp.Phone,
			Email:     resp.Email,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			DeletedAt: timestamppb.New(resp.DeletedAt.Time),
		},
		Car: &obdv1car.CarView{
			CarId:           int32(resp.CarID.Int64),
			CarBrandId:      int32(resp.CarBrandID.Int64),
			CarBrandModelId: int32(resp.CarBrandModelID.Int64),
			BrandName:       resp.BrandName.String,
			BrandModelName:  resp.BrandModelName.String,
			ModelYear:       resp.ModelYear.Int32,
		},
	}
	return serverResp, nil
}
