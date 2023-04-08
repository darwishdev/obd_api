package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
)

func (u *CenterUsecase) CentersList(ctx context.Context, req *obdv1.CentersListRequest) (*obdv1.CentersListResponse, error) {
	centers, err := u.repo.CentersList(ctx, &req.AreaId)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.NewCentersListFromSqlResponse(centers)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
