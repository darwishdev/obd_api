package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) SessionsList(ctx context.Context, req *connect.Request[obdv1.SessionsListRequest]) (*connect.Response[obdv1.SessionsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	authPayload, err := server.authorizeUser(req.Header())
	if err != nil {
		return nil, validator.UnauthErr(err)
	}

	resp, err := server.sessionUsecase.SessionsList(ctx, &authPayload.CarId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
