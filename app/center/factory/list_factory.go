package factory

import (
	"math"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func centersListRequestValidation(req *obdv1.CentersListRequest) error {
	if err := validator.ValidateLatLng(req.Lat, req.Long); err != nil {
		return validator.InvalidArgErr(err)
	}
	return nil
}
func (f *CenterFactory) ListSqlFromGrpc(req *obdv1.CentersListRequest) (*db.CentersListParams, error) {
	if err := centersListRequestValidation(req); err != nil {
		return nil, err
	}
	return &db.CentersListParams{
		Lat:  float64(req.Lat),
		Long: float64(req.Long),
	}, nil
}
func (f *CenterFactory) ListGrpcFromSql(req *[]db.CenterInfo) (*obdv1.CentersListResponse, error) {
	centers := make([]*obdv1.CenterListRow, 0, len(*req))
	for _, center := range *req {
		c, err := NewCenterFromSqlResponse(&center)
		if err != nil {
			return nil, err
		}
		centers = append(centers, c)
	}
	return &obdv1.CentersListResponse{
		Centers: centers,
	}, nil
}

func NewCenterFromSqlResponse(req *db.CenterInfo) (*obdv1.CenterListRow, error) {
	response := &obdv1.CenterListRow{
		CenterId:  int32(req.CenterID),
		Name:      req.Name,
		Phone:     req.Phone,
		Image:     req.Image,
		Address:   req.Address,
		AreaId:    int32(req.AreaID),
		Lat:       req.Lat,
		Distance:  req.Distance,
		Long:      req.Long,
		Rate:      int32(math.Round(req.AvgRate)),
		CreatedAt: timestamppb.New(req.CreatedAt),
		DeletedAt: nil,
	}

	log.Debug().Interface("interface", req.AvgRate).Msg("deleteme")
	return response, nil
}
