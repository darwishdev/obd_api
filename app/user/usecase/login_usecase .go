package usecase

import (
	"context"
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/darwishdev/obd_api/pkg/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserUsecase) generateUserToken(email string, userID int64) *obdv1.LoginInfo {
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(
		email,
		userID,
		u.config.AccessTokenDuration,
	)
	if err != nil {
		return nil
	}

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(
		email,
		userID,
		u.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil
	}

	return &obdv1.LoginInfo{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
}
func (u *UserUsecase) UserLogin(ctx context.Context, req *obdv1.UserLoginRequest) (*obdv1.UserLoginResponse, error) {
	err := u.factory.UserLoginRequestValidation(req)
	if err != nil {
		return nil, err
	}

	record, err := u.repo.UserGetByUsername(ctx, &req.Email)
	if err != nil {
		return nil, err
	}
	err = util.CheckPassword(req.Password, record.Password)
	if err != nil {
		return nil, validator.InvalidArgErr(fmt.Errorf("incorrect_password"))
	}

	resp, err := u.factory.LoginGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	resp.LoginInfo = u.generateUserToken(record.Email, record.UserID)
	return resp, nil
}
