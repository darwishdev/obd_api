package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *SessionFactory) CloseGrpcFromSql(req *db.Session) (*obdv1.SessionCloseResponse, error) {
	response := &obdv1.SessionCloseResponse{
		Session: &obdv1.Session{
			SessionId:  int32(req.SessionID),
			CarId:      int32(req.CarID),
			IsLive:     req.IsLive,
			CreatedAt:  timestamppb.New(req.CreatedAt),
			FinishedAt: timestamppb.New(req.FinishedAt.Time),
		},
	}

	return response, nil
}
