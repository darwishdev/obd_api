package factory

import (
	"database/sql"
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func reviewsListRequestValidation(req *obdv1.ReviewsListRequest) error {
	if err := validator.ValidateUnsignedInt(req.CenterId); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("center_id_%w", err))
	}
	return nil
}
func listGrpcFromSql(req *db.ReviewsListRow) (*obdv1.ReviewListRow, error) {
	response := &obdv1.ReviewListRow{
		UserName:   req.UserName,
		CenterName: req.CenterName,
		Review:     req.Review,
		Rate:       int32(req.Rate),
		CreatedAt:  timestamppb.New(req.CreatedAt),
	}

	return response, nil
}
func (f *ReviewFactory) ListSqlFromGrpc(req *obdv1.ReviewsListRequest) (*db.ReviewsListParams, error) {
	if err := reviewsListRequestValidation(req); err != nil {
		return nil, err
	}
	var resp db.ReviewsListParams

	resp.CenterID = sql.NullInt64{Int64: req.CenterId, Valid: req.CenterId != 0}

	log.Debug().Interface("resp", resp).Msg("Dekete")
	return &resp, nil
}

func (f *ReviewFactory) ListGrpcFromSqlArr(req *[]db.ReviewsListRow) (*obdv1.ReviewsListResponse, error) {
	reviews := make([]*obdv1.ReviewListRow, 0, len(*req))
	for _, review := range *req {
		c, err := listGrpcFromSql(&review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, c)
	}
	return &obdv1.ReviewsListResponse{
		Reviews: reviews,
	}, nil
}
