package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type SessionRepoInterface interface {
	SessionCreate(ctx context.Context, req *db.SessionCreateParams) (*db.Session, error)
	SessionAttachCodes(ctx context.Context, req *db.SessionAttachCodesParams) (*[]db.SessionCode, error)
	SessionAttachValues(ctx context.Context, req *db.SessionAttachValuesParams) (*[]db.SessionValue, error)
	SessionsList(ctx context.Context, req *int64) (*[]db.SessionsListRow, error)
	SessionGetCodes(ctx context.Context, req *int64) (*[]db.Code, error)
	SessionClose(ctx context.Context, req *int64) (*db.Session, error)
}

type SessionRepo struct {
	store db.Store
}

func NewSessionRepo(store db.Store) *SessionRepo {
	return &SessionRepo{
		store: store,
	}
}
