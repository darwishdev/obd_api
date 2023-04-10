package usecase

import (
	"context"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
)

func (u *CarUsecase) BrandsList(ctx context.Context, req *obdv1.BrandsListRequest) (*obdv1.BrandsListResponse, error) {
	brands, err := u.repo.BrandsList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := u.factory.BrandsListGrpcFromSqlArr(brands)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
