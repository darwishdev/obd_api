package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
)

func (u *UserUsecase) UserAuthorize(ctx context.Context, req int64) (*obdv1.UserAuthorizeResponse, error) {
	record, err := u.repo.UserGet(ctx, &req)
	if err != nil {
		return nil, err
	}
	resp, err := u.factory.AuthorizeGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	resp.LoginInfo = u.generateUserToken(record.Email, record.UserID, record.CarID.Int64)
	return resp, nil
}
