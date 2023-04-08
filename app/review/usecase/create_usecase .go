package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
)

func (u *ReviewUsecase) ReviewCreate(ctx context.Context, req *obdv1.ReviewCreateRequest, userId int64) (*obdv1.ReviewCreateResponse, error) {
	request, err := u.factory.CreateSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}

	request.UserID = userId
	reviews, err := u.repo.ReviewCreate(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.CreateGrpcFromSql(reviews)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
