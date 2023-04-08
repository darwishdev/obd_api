package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"

	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) UserLogin(ctx context.Context, req *connect.Request[obdv1.UserLoginRequest]) (*connect.Response[obdv1.UserLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	resp, err := server.userUsecase.UserLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
