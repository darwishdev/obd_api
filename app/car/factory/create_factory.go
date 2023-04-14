package factory

import (
	"fmt"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func carCreateRequestValidation(req *obdv1.CarCreateRequest) error {
	if err := validator.ValidateUnsignedInt(req.CarBrandModelId); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("car_brand_model_id_%w", err))
	}
	if err := validator.ValidateInt(int64(req.ModelYear), 1990, int64(time.Now().Year())); err != nil {
		return validator.InvalidArgErr(fmt.Errorf("model_year_%w", err))
	}
	return nil
}
func (f *CarFactory) CreateSqlFromGrpc(req *obdv1.CarCreateRequest) (*db.CarCreateParams, error) {
	if err := carCreateRequestValidation(req); err != nil {
		return nil, err
	}
	resp := &db.CarCreateParams{
		CarBrandModelID: req.CarBrandModelId,
		ModelYear:       req.ModelYear,
	}
	return resp, nil
}

func (f *CarFactory) CreateGrpcFromSql(req *db.Car) (*obdv1.CarCreateResponse, error) {
	response := &obdv1.CarCreateResponse{
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
