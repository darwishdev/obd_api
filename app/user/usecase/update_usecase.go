package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	"github.com/rs/zerolog/log"
)

func (u *UserUsecase) UserUpdate(ctx context.Context, req *obdv1.UserUpdateRequest, userId int64) (*obdv1.UserUpdateResponse, error) {
	log.Logger.Debug().Interface("req", req).Msg("remove me")
	s, err := u.factory.NewUserUpdateFromProto(req)
	if err != nil {
		return nil, err
	}
	s.UserID = userId
	record, err := u.repo.UserUpdate(ctx, s)
	if err != nil {
		return nil, err
	}
	res, err := u.factory.NewUserUpdateFromSqlResponse(record)
	if err != nil {
		return nil, err
	}
	return res, nil
}
