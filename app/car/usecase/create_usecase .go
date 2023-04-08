package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
)

func (u *CarUsecase) CarCreate(ctx context.Context, req *obdv1.CarCreateRequest, userId int64) (*obdv1.CarCreateResponse, error) {
	request, err := u.factory.CreateSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}

	request.UserID = userId
	cars, err := u.repo.CarCreate(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.CreateGrpcFromSql(cars)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
