package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func winchListRequestValidation(req *obdv1.WinchListRequest) error {
	if err := validator.ValidateLatLng(req.Lat, req.Long); err != nil {
		return validator.InvalidArgErr(err)
	}
	return nil
}
func (f *WinchFactory) ListSqlFromGrpc(req *obdv1.WinchListRequest) (*db.WinchListParams, error) {
	if err := winchListRequestValidation(req); err != nil {
		return nil, err
	}
	resp := &db.WinchListParams{
		InLat:  float64(req.Lat),
		InLong: float64(req.Long),
	}
	return resp, nil
}
func listGrpcFromSql(req *db.WinchInfo) (*obdv1.WinchListRow, error) {
	response := &obdv1.WinchListRow{
		WinchID:     int32(req.WinchID),
		Name:        req.Name,
		Phone:       req.Phone,
		DriverName:  req.DriverName,
		DriverPhone: req.DriverPhone,
		Lat:         req.Long,
		Long:        req.Lat,
		Distance:    req.Distance,
		CreatedAt:   timestamppb.New(req.CreatedAt),
	}

	return response, nil
}

func (f *WinchFactory) ListGrpcFromSqlArr(req *[]db.WinchInfo) (*obdv1.WinchListResponse, error) {
	winches := make([]*obdv1.WinchListRow, 0, len(*req))
	for _, winch := range *req {
		c, err := listGrpcFromSql(&winch)
		if err != nil {
			return nil, err
		}
		winches = append(winches, c)
	}
	return &obdv1.WinchListResponse{
		Winches: winches,
	}, nil
}
