package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	"github.com/darwishdev/obd_api/pkg/validator"
	"github.com/rs/zerolog/log"
)

func (server *Server) CentersList(ctx context.Context, req *connect.Request[obdv1.CentersListRequest]) (*connect.Response[obdv1.CentersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}

	resp, err := server.centerUsecase.CentersList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(resp)

	log.Debug().Interface("res", res.Msg.Centers[0].CenterId).Msg("test rest")
	return res, nil
}
