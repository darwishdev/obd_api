package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
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
		CenterId:  req.CenterID,
		Name:      req.Name,
		Phone:     req.Phone,
		Location:  req.Location,
		Address:   req.Address,
		AreaId:    req.AreaID,
		Lat:       req.Lat,
		Distance:  req.Distance,
		Long:      req.Long,
		CreatedAt: timestamppb.New(req.CreatedAt),
		DeletedAt: nil,
	}

	return response, nil
}
