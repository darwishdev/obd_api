package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) SessionClose(ctx context.Context, req *connect.Request[obdv1.SessionCloseRequest]) (*connect.Response[obdv1.SessionCloseResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	_, err := server.authorizeUser(req.Header())
	if err != nil {
		return nil, validator.UnauthErr(err)
	}
	resp, err := server.sessionUsecase.SessionClose(ctx, &req.Msg.SessionId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
