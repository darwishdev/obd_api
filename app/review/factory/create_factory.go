package factory

import (
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func reviewCreateRequestValidation(req *obdv1.ReviewCreateRequest) error {
	if err := validator.ValidateUnsignedInt(req.CenterId); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("center_id_%w", err))
	}
	if err := validator.ValidateInt(int64(req.Rate), 1, 5); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("review_id_%w", err))
	}
	if err := validator.ValidateString(req.Review, 3, 200); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("review_%w", err))
	}
	return nil
}
func (f *ReviewFactory) CreateSqlFromGrpc(req *obdv1.ReviewCreateRequest) (*db.ReviewCreateParams, error) {
	if err := reviewCreateRequestValidation(req); err != nil {
		return nil, err
	}
	resp := &db.ReviewCreateParams{
		CenterID: req.CenterId,
		Review:   req.Review,
		Rate:     int16(req.Rate),
	}
	return resp, nil
}

func (f *ReviewFactory) CreateGrpcFromSql(req *db.Review) (*obdv1.ReviewCreateResponse, error) {
	response := &obdv1.ReviewCreateResponse{
		Review: &obdv1.Review{
			ReviewId:  int32(req.ReviewID),
			UserId:    int32(req.UserID),
			CenterId:  int32(req.CenterID),
			Review:    req.Review,
			Rate:      int32(req.Rate),
			CreatedAt: timestamppb.New(req.CreatedAt),
		},
	}

	return response, nil
}
