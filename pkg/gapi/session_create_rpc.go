package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) SessionCreate(ctx context.Context, req *connect.Request[obdv1.SessionCreateRequest]) (*connect.Response[obdv1.SessionCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	authPayload, err := server.authorizeUser(req.Header())
	if err != nil {
		return nil, validator.UnauthErr(err)
	}
	resp, err := server.sessionUsecase.SessionCreate(ctx, req.Msg, authPayload.CarId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *Server) SessionAttachCode(ctx context.Context, req *connect.Request[obdv1.SessionAttachCodeRequest]) (*connect.Response[obdv1.SessionAttachCodeResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.sessionUsecase.SessionAttachCode(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *Server) SessionAttachValue(ctx context.Context, req *connect.Request[obdv1.SessionAttachValueRequest]) (*connect.Response[obdv1.SessionAttachValueResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.sessionUsecase.SessionAttachValue(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
