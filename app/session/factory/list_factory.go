package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func listGrpcFromSql(req *db.SessionsListRow) (*obdv1.Session, error) {
	response := &obdv1.Session{
		SessionId:    int32(req.SessionID),
		CarId:        int32(req.CarID),
		IsLive:       req.IsLive,
		Emergencies:  int32(req.Emergencies),
		DefaultCodes: int32(req.DefaultCodes),
		CreatedAt:    timestamppb.New(req.CreatedAt),
		FinishedAt:   timestamppb.New(req.FinishedAt.Time),
	}

	return response, nil
}

func (f *SessionFactory) ListGrpcFromSqlArr(req *[]db.SessionsListRow) (*obdv1.SessionsListResponse, error) {
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

func getCodesGrpcFromSql(req *db.Code) (*obdv1.Code, error) {
	response := &obdv1.Code{
		CodeId:          int32(req.CodeID),
		CarBrandModelId: int32(req.CarBrandModelID),
		CodeName:        req.CodeName,
		VehiclePart:     req.VehiclePart,
		CodeType:        req.CodeType,
		Description:     req.Description,
		IsEmergency:     req.IsEmergency.Bool,
	}

	return response, nil
}

func (f *SessionFactory) GetCodesGrpcFromSqlArr(req *[]db.Code) (*obdv1.SessionGetCodesResponse, error) {
	codes := make([]*obdv1.Code, 0, len(*req))
	log.Debug().Int("req", len(*req)).Msg("delete  me")
	for _, session := range *req {
		c, err := getCodesGrpcFromSql(&session)
		if err != nil {
			return nil, err
		}
		codes = append(codes, c)
	}
	return &obdv1.SessionGetCodesResponse{
		Codes: codes,
	}, nil
}
