package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) CarBrandsList(ctx context.Context, req *connect.Request[obdv1.BrandsListRequest]) (*connect.Response[obdv1.BrandsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	resp, err := server.carUsecase.BrandsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
