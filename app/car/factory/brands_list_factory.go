package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/rs/zerolog/log"
)

func brandModelGrpcFromSql(req *db.CarBrandsListRow) *obdv1.CarBrandModel {
	return &obdv1.CarBrandModel{CarBrandModelId: req.CarBrandModelID, Name: req.ModelName, Years: req.Years}
}
func (f *CarFactory) BrandsListGrpcFromSqlArr(req *[]db.CarBrandsListRow) (*obdv1.BrandsListResponse, error) {
	var response []*obdv1.BrandsListRow
	var currentBrnad string
	for _, rec := range *req {
		if rec.Name != currentBrnad {
			currentBrnad = rec.Name
			response = append(response, &obdv1.BrandsListRow{Name: rec.Name, Models: []*obdv1.CarBrandModel{brandModelGrpcFromSql(&rec)}})
			log.Debug().Interface("response", response).Msg("reponse interface")
		} else {
			response[len(response)-1].Models = append(response[len(response)-1].Models, brandModelGrpcFromSql(&rec))
		}
	}
	return &obdv1.BrandsListResponse{Brands: response}, nil
}
