package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"

	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) UserCreate(ctx context.Context, req *connect.Request[obdv1.UserCreateRequest]) (*connect.Response[obdv1.UserCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.userUsecase.UserCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
