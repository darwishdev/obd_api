package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type ReviewFactoryInterface interface {
	ListSqlFromGrpc(req *obdv1.ReviewsListRequest) (*db.ReviewsListParams, error)
	ListGrpcFromSqlArr(req *[]db.ReviewsListRow) (*obdv1.ReviewsListResponse, error)
	CreateSqlFromGrpc(req *obdv1.ReviewCreateRequest) (*db.ReviewCreateParams, error)
	CreateGrpcFromSql(req *db.Review) (*obdv1.ReviewCreateResponse, error)
}

type ReviewFactory struct {
}

func NewReviewFactory() *ReviewFactory {
	return &ReviewFactory{}
}
