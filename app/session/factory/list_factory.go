package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func listGrpcFromSql(req *db.Session) (*obdv1.Session, error) {
	response := &obdv1.Session{
		SessionId:  int32(req.SessionID),
		CarId:      int32(req.CarID),
		IsLive:     req.IsLive,
		CreatedAt:  timestamppb.New(req.CreatedAt),
		FinishedAt: timestamppb.New(req.FinishedAt.Time),
	}

	return response, nil
}

func (f *SessionFactory) ListGrpcFromSqlArr(req *[]db.Session) (*obdv1.SessionsListResponse, error) {
	sessions := make([]*obdv1.Session, 0, len(*req))
	for _, session := range *req {
		c, err := listGrpcFromSql(&session)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, c)
	}
	return &obdv1.SessionsListResponse{
		Sessions: sessions,
	}, nil
}
