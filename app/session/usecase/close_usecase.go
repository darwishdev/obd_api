package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
)

func (u *SessionUsecase) SessionClose(ctx context.Context, sessionId *int64) (*obdv1.SessionCloseResponse, error) {

	sessoion, err := u.repo.SessionClose(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	resp, err := u.factory.CloseGrpcFromSql(sessoion)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
