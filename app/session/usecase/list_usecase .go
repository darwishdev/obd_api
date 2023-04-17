package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
)

func (u *SessionUsecase) SessionsList(ctx context.Context, carId *int64) (*obdv1.SessionsListResponse, error) {

	sessoions, err := u.repo.SessionsList(ctx, carId)
	if err != nil {
		return nil, err
	}
	resp, err := u.factory.ListGrpcFromSqlArr(sessoions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
