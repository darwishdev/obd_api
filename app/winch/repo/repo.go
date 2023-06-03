package repo

import (
	"context"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
)

type WinchRepoInterface interface {
	WinchList(ctx context.Context, req *db.WinchListParams) (*[]db.WinchInfo, error)
}

type WinchRepo struct {
	store db.Store
}

func NewWinchRepo(store db.Store) *WinchRepo {
	return &WinchRepo{
		store: store,
	}
}
