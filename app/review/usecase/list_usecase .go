package usecase

import (
	"context"
	"database/sql"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
)

func (u *ReviewUsecase) ReviewsList(ctx context.Context, req *obdv1.ReviewsListRequest, userId int64) (*obdv1.ReviewsListResponse, error) {
	request, err := u.factory.ListSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	// ceheck if center not passed to pass the user id
	request.UserID = sql.NullInt64{Int64: userId, Valid: !request.CenterID.Valid && userId != 0}
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
