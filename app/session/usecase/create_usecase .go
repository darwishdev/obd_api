package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

func (u *SessionUsecase) SessionCreate(ctx context.Context, req *obdv1.SessionCreateRequest, carId int64) (*obdv1.SessionCreateResponse, error) {
	request := &db.SessionCreateParams{
		CarID:  carId,
		IsLive: req.IsLive,
	}
	sessions, err := u.repo.SessionCreate(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.CreateGrpcFromSql(sessions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *SessionUsecase) SessionAttachCodes(ctx context.Context, req *obdv1.SessionAttachCodesRequest) (*obdv1.SessionAttachCodesResponse, error) {
	request, err := u.factory.AttachCodesSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	_, err = u.repo.SessionAttachCodes(ctx, request)
	if err != nil {
		return nil, err
	}

	resp := &obdv1.SessionAttachCodesResponse{
		Created: true,
	}

	return resp, nil
}

func (u *SessionUsecase) SessionAttachValues(ctx context.Context, req *obdv1.SessionAttachValuesRequest) (*obdv1.SessionAttachValuesResponse, error) {
	request, err := u.factory.AttachValuesSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	_, err = u.repo.SessionAttachValues(ctx, request)
	if err != nil {
		return nil, err
	}

	resp := &obdv1.SessionAttachValuesResponse{
		Created: true,
	}

	return resp, nil
}
