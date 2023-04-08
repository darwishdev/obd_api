package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCenterFromSqlResponse(req *db.Center) (*obdv1.Center, error) {
	response := &obdv1.Center{
		CenterId:  req.CenterID,
		Name:      req.Name,
		Phone:     req.Phone,
		Location:  req.Location,
		Address:   req.Address,
		AreaId:    req.AreaID,
		Lat:       req.Lat,
		Long:      req.Long,
		CreatedAt: timestamppb.New(req.CreatedAt),
		DeletedAt: nil,
	}

	return response, nil
}

func (f *CenterFactory) NewCentersListFromSqlResponse(req *[]db.Center) (*obdv1.CentersListResponse, error) {
	centers := make([]*obdv1.Center, 0, len(*req))
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
