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

func (u *SessionUsecase) SessionAttachCode(ctx context.Context, req *obdv1.SessionAttachCodeRequest) (*obdv1.SessionAttachCodeResponse, error) {
	request, err := u.factory.AttachCodeSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	sessions, err := u.repo.SessionAttachCode(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.AttachCodeGrpcFromSql(sessions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *SessionUsecase) SessionAttachValue(ctx context.Context, req *obdv1.SessionAttachValueRequest) (*obdv1.SessionAttachValueResponse, error) {
	request, err := u.factory.AttachValueSqlFromGrpc(req)
	if err != nil {
		return nil, err
	}
	sessions, err := u.repo.SessionAttachValue(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, err := u.factory.AttachValueGrpcFromSql(sessions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
