package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) WinchList(ctx context.Context, req *connect.Request[obdv1.WinchListRequest]) (*connect.Response[obdv1.WinchListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	resp, err := server.winchUsecase.WinchList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
