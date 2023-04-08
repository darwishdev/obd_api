package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) UserAuthorize(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[obdv1.UserAuthorizeResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	authPayload, err := server.authorizeUser(req.Header())
	if err != nil {
		return nil, validator.UnauthErr(err)
	}
	resp, err := server.userUsecase.UserAuthorize(ctx, authPayload.UserId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
