package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
)

func (u *WinchUsecase) WinchList(ctx context.Context, req *obdv1.WinchListRequest) (*obdv1.WinchListResponse, error) {
	request, err := u.factory.ListSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	winches, err := u.repo.WinchList(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.ListGrpcFromSqlArr(winches)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
