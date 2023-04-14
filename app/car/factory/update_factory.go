package factory

import (
	"database/sql"
	"fmt"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func carUpdateRequestValidation(req *obdv1.CarUpdateRequest) error {
	if req.CarBrandModelId != 0 {
		if err := validator.ValidateUnsignedInt(req.CarBrandModelId); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("car_brand_model_id_%w", err))
		}
	}
	if req.ModelYear != 0 {
		if err := validator.ValidateInt(int64(req.ModelYear), 1990, int64(time.Now().Year())); err != nil {
			return validator.InvalidArgErr(fmt.Errorf("model_year_%w", err))
		}
	}
	return nil
}
func (f *CarFactory) UpdateSqlFromGrpc(req *obdv1.CarUpdateRequest) (*db.CarUpdateParams, error) {
	if err := carUpdateRequestValidation(req); err != nil {
		return nil, err
	}
	resp := &db.CarUpdateParams{
		CarID:           req.CarId,
		CarBrandModelID: sql.NullInt64{Int64: req.CarBrandModelId},
		ModelYear:       sql.NullInt32{Int32: req.ModelYear},
	}

	log.Debug().Interface("resp", resp)
	return resp, nil
}

func (f *CarFactory) UpdateGrpcFromSql(req *db.Car) (*obdv1.CarUpdateResponse, error) {
	response := &obdv1.CarUpdateResponse{
		Car: &obdv1.Car{
			CarId:           int32(req.CarID),
			CarBrandModelId: int32(req.CarBrandModelID),
			UserId:          int32(req.UserID),
			ModelYear:       req.ModelYear,
			CreatedAt:       timestamppb.New(req.CreatedAt),
		},
	}

	return response, nil
}
