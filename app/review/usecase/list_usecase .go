package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
)

func (u *ReviewUsecase) ReviewsList(ctx context.Context, req *obdv1.ReviewsListRequest) (*obdv1.ReviewsListResponse, error) {
	request, err := u.factory.ListSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	reviews, err := u.repo.ReviewsList(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.ListGrpcFromSqlArr(reviews)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
