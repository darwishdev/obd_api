package factory

import (
	"reflect"
	"strings"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/rs/zerolog/log"
)

func (f *SessionFactory) CreateGrpcFromSql(req *db.Session) (*obdv1.SessionCreateResponse, error) {
	response := &obdv1.SessionCreateResponse{
		SessionId: int32(req.SessionID),
	}

	return response, nil
}

func (f *SessionFactory) AttachCodesSqlFromGrpc(req *obdv1.SessionAttachCodesRequest) (*db.SessionAttachCodesParams, error) {
	request := &db.SessionAttachCodesParams{
		SessionID:    req.SessionId,
		SessionCodes: req.Codes,
	}

	return request, nil
}

func _extractOBDModel(metrics *obdv1.OBDModel) (string, string) {
	var keys []string
	var values []string
	v := reflect.ValueOf(*metrics)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() != "" {
			keys = append(keys, v.Type().Field(i).Name)
			values = append(values, field.String())
		}
	}
	keyString := strings.Join(keys, ",")
	valueString := strings.Join(values, ",")
	return keyString, valueString
}
func (f *SessionFactory) AttachValuesSqlFromGrpc(req *obdv1.SessionAttachValuesRequest) (*db.SessionAttachValuesParams, error) {
	keys, values := _extractOBDModel(req.ObdModel)

	log.Info().
		Str("keys", keys).
		Str("values", values).
		Msg("delete me")
	request := &db.SessionAttachValuesParams{
		SessionID:     req.SessionId,
		SessionKeys:   keys,
		SessionValues: values,
	}

	return request, nil
}
