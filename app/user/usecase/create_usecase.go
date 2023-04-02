package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
)

func (u *UserUsecase) UserCreate(ctx context.Context, req *obdv1.UserCreateRequest) (*obdv1.UserCreateResponse, error) {
	s, err := u.factory.NewUserFromProto(req)
	if err != nil {
		return nil, err
	}
	record, err := u.repo.UserCreate(ctx, s)
	if err != nil {
		return nil, err
	}
	res, err := u.factory.NewUserCreateFromSqlResponse(record)
	if err != nil {
		return nil, err
	}
	return res, nil
}
