package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type SessionRepoInterface interface {
	SessionCreate(ctx context.Context, req *db.SessionCreateParams) (*db.Session, error)
	SessionAttachCode(ctx context.Context, req *db.SessionAttachCodeParams) (*db.SessionCode, error)
	SessionAttachValue(ctx context.Context, req *db.SessionAttachValueParams) (*db.SessionValue, error)
}

type SessionRepo struct {
	store db.Store
}

func NewSessionRepo(store db.Store) *SessionRepo {
	return &SessionRepo{
		store: store,
	}
}
