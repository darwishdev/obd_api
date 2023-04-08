package usecase

import (
	"context"
	"fmt"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserUsecase) UserLogin(ctx context.Context, req *obdv1.UserLoginRequest) (*obdv1.UserLoginResponse, error) {
	err := u.factory.UserLoginRequestValidation(req)
	if err != nil {
		return nil, err
	}

	record, err := u.repo.UserGetByUsername(ctx, &req.Email)
	if err != nil {
		return nil, err
	}

	log.Logger.Debug().Str("password", record.Password).Msg("user password")
	log.Logger.Debug().Str("password", req.Password).Msg("request password")
	err = util.CheckPassword(req.Password, record.Password)
	if err != nil {
		return nil, validator.InvalidArgErr(fmt.Errorf("incorrect_password"))
	}
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(
		record.Email,
		record.UserID,
		u.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(
		record.Email,
		record.UserID,
		u.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	resp, err := u.factory.LoginGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	resp.AccessToken = accessToken
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	resp.AccessTokenExpiresAt = timestamppb.New(accessPayload.ExpiredAt)
	resp.RefreshTokenExpiresAt = timestamppb.New(refreshPayload.ExpiredAt)
	return resp, nil
}
