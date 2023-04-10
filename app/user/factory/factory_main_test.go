package factory

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"
	"github.com/darwishdev/obd_api/pkg/validator"
	_ "github.com/lib/pq"
)

var (
	factory               UserFactoryInterface
	validUserResponse     *db.UserInfo
	ErrInvalidName        error
	ErrInvalidPassword    error
	ErrInvalidEmailLength error
	ErrInvalidEmail       error
	ErrInvalidPhone       error
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewUserFactory()
	validUserResponse = &db.UserInfo{
		Name:           util.RandomName(),
		Email:          util.RandomEmail(),
		Phone:          util.RandomPhone(),
		Password:       util.RandomString(6),
		BrandName:      sql.NullString{String: "Toyota"},
		ModelYear:      sql.NullInt32{Int32: 2020},
		BrandModelName: sql.NullString{String: "Corolla"},
		CreatedAt:      time.Now(),
	}
	ErrInvalidName = validator.InvalidArgErr(fmt.Errorf("name_%w", validator.ErrorMinMax(float32(3), float32(200))))
	ErrInvalidPassword = validator.InvalidArgErr(fmt.Errorf("password_%w", validator.ErrorMinMax(float32(6), float32(200))))
	ErrInvalidEmailLength = validator.InvalidArgErr(fmt.Errorf("email_%w", validator.ErrorMinMax(float32(3), float32(200))))
	ErrInvalidEmail = validator.InvalidArgErr(fmt.Errorf("email_%w", validator.ErrorInvalid()))
	ErrInvalidPhone = validator.InvalidArgErr(fmt.Errorf("phone_%w", validator.ErrorInvalid()))
	// Cleanup, if necessary

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
