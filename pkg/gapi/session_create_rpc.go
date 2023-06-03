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

func (server *Server) SessionAttachCodes(ctx context.Context, req *connect.Request[obdv1.SessionAttachCodesRequest]) (*connect.Response[obdv1.SessionAttachCodesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.sessionUsecase.SessionAttachCodes(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *Server) SessionAttachValues(ctx context.Context, req *connect.Request[obdv1.SessionAttachValuesRequest]) (*connect.Response[obdv1.SessionAttachValuesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.sessionUsecase.SessionAttachValues(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
