package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) ReviewsList(ctx context.Context, req *connect.Request[obdv1.ReviewsListRequest]) (*connect.Response[obdv1.ReviewsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	resp, err := server.reviewUsecase.ReviewsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
