package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) CarCreate(ctx context.Context, req *connect.Request[obdv1.CarCreateRequest]) (*connect.Response[obdv1.CarCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	authPayload, err := server.authorizeUser(req.Header())
	if err != nil {
		return nil, validator.UnauthErr(err)
	}
	resp, err := server.carUsecase.CarCreate(ctx, req.Msg, authPayload.UserId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
