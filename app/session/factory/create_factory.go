package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *SessionFactory) CreateGrpcFromSql(req *db.Session) (*obdv1.SessionCreateResponse, error) {
	response := &obdv1.SessionCreateResponse{
		Session: &obdv1.Session{
			SessionId: int32(req.SessionID),
			CarId:     int32(req.CarID),
			IsLive:    req.IsLive,
			CreatedAt: timestamppb.New(req.CreatedAt),
		},
	}

	return response, nil
}

func (f *SessionFactory) AttachCodeSqlFromGrpc(req *obdv1.SessionAttachCodeRequest) (*db.SessionAttachCodeParams, error) {
	request := &db.SessionAttachCodeParams{
		SessionID: req.SessionId,
		CodeID:    req.CodeId,
	}

	return request, nil
}
func (f *SessionFactory) AttachCodeGrpcFromSql(resp *db.SessionCode) (*obdv1.SessionAttachCodeResponse, error) {
	response := &obdv1.SessionAttachCodeResponse{
		Session: &obdv1.SessionCode{
			SessionCodeId: int32(resp.SessionCodeID),
			SessionId:     int32(resp.SessionID),
			CodeId:        int32(resp.CodeID),
		},
	}

	return response, nil
}

func (f *SessionFactory) AttachValueSqlFromGrpc(req *obdv1.SessionAttachValueRequest) (*db.SessionAttachValueParams, error) {
	request := &db.SessionAttachValueParams{
		SessionID: req.SessionId,
		ValueKey:  req.ValueKey,
		ValueData: req.ValueData,
	}

	return request, nil
}
func (f *SessionFactory) AttachValueGrpcFromSql(resp *db.SessionValue) (*obdv1.SessionAttachValueResponse, error) {
	response := &obdv1.SessionAttachValueResponse{
		Session: &obdv1.SessionValue{
			SessionValueId: int32(resp.SessionValueID),
			SessionId:      int32(resp.SessionID),
			ValueKey:       resp.ValueKey,
			ValueData:      resp.ValueData,
		},
	}

	return response, nil
}
