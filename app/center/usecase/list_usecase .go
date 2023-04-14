package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
)

func (u *CenterUsecase) CentersList(ctx context.Context, req *obdv1.CentersListRequest) (*obdv1.CentersListResponse, error) {
	request, err := u.factory.ListSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	centers, err := u.repo.CentersList(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.ListGrpcFromSql(centers)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
