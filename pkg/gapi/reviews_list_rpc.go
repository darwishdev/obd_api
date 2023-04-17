package gapi

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	"github.com/darwishdev/obd_api/pkg/validator"
)

func (server *Server) ReviewsList(ctx context.Context, req *connect.Request[obdv1.ReviewsListRequest]) (*connect.Response[obdv1.ReviewsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, validator.InternalErr(err)
	}
	// try to authorize the user to check if token passed (means err == nil) to pass value of user id to usecae
	// and if if there is error we will do nothing (means send 0 to userId)
	var userId int64
	authPayload, err := server.authorizeUser(req.Header())
	if err != nil {
		if req.Msg.CenterId == 0 {
			return nil, validator.InvalidArgErr(fmt.Errorf("pass_center_id_or_user_token"))
		}
	} else {
		userId = authPayload.UserId
	}

	resp, err := server.reviewUsecase.ReviewsList(ctx, req.Msg, userId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
