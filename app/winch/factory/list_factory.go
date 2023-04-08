package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func listGrpcFromSql(req *db.WinchListRow) (*obdv1.Winch, error) {
	response := &obdv1.Winch{
		WinchId:     req.WinchID,
		Name:        req.Name,
		Phone:       req.Phone,
		DriverName:  req.DriverName,
		DriverPhone: req.DriverPhone,
		CreatedAt:   timestamppb.New(req.CreatedAt),
		DeletedAt:   nil,
	}

	return response, nil
}

func (f *WinchFactory) ListGrpcFromSqlArr(req *[]db.WinchListRow) (*obdv1.WinchListResponse, error) {
	winches := make([]*obdv1.Winch, 0, len(*req))
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
