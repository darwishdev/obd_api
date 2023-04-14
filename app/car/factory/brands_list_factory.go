package factory

import (
	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/rs/zerolog/log"
)

func brandModelGrpcFromSql(req *db.CarBrandsListRow) *obdv1.CarBrandModel {
	return &obdv1.CarBrandModel{CarBrandModelId: int32(req.CarBrandModelID), Name: req.ModelName, Years: req.Years}
}
func (f *CarFactory) BrandsListGrpcFromSqlArr(req *[]db.CarBrandsListRow) (*obdv1.BrandsListResponse, error) {
	var response []*obdv1.BrandsListRow
	var currentBrnadName string
	var currentBrnadId int32
	for _, rec := range *req {
		if int32(rec.CarBrandID) != currentBrnadId {
			currentBrnadName = rec.Name
			currentBrnadId = int32(rec.CarBrandID)
			response = append(response, &obdv1.BrandsListRow{Name: currentBrnadName, CarBrandId: currentBrnadId, Models: []*obdv1.CarBrandModel{brandModelGrpcFromSql(&rec)}})
			log.Debug().Interface("response", response).Msg("reponse interface")
		} else {
			response[len(response)-1].Models = append(response[len(response)-1].Models, brandModelGrpcFromSql(&rec))
		}
	}
	return &obdv1.BrandsListResponse{Brands: response}, nil
}
