package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
)

func (u *CarUsecase) CarUpdate(ctx context.Context, req *obdv1.CarUpdateRequest, userId int64) (*obdv1.CarUpdateResponse, error) {
	request, err := u.factory.UpdateSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}

	cars, err := u.repo.CarUpdate(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.UpdateGrpcFromSql(cars)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
