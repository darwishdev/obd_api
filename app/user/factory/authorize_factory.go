package factory

import (
	obdv1Car "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	obdv1User "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *UserFactory) AuthorizeGrpcFromSql(resp *db.UserInfo) (*obdv1User.UserAuthorizeResponse, error) {
	serverResp := &obdv1User.UserAuthorizeResponse{
		User: &obdv1User.User{
			UserId:            resp.UserID,
			Name:              resp.Name,
			Phone:             resp.Phone,
			Email:             resp.Email,
			Password:          resp.Password,
			PasswordChangedAt: timestamppb.New(resp.PasswordChangedAt),
			CreatedAt:         timestamppb.New(resp.CreatedAt),
		},
	}
	if resp.BrandName.Valid {
		carInfo := &obdv1Car.CarView{
			BrandName:      resp.BrandName.String,
			BrandModelName: resp.BrandModelName.String,
			ModelYear:      resp.ModelYear.Int32,
		}
		serverResp.Car = carInfo
	}

	return serverResp, nil
}
