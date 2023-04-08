package repo

import (
	"os"
	"testing"

	mockdb "github.com/darwishdev/obd_api/pkg/sqlc/mock"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

var (
	store *mockdb.MockStore
	repo  CenterRepoInterface
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	ctrl := gomock.NewController(&testing.T{})
	defer ctrl.Finish()
	store = mockdb.NewMockStore(ctrl)
	repo = NewCenterRepo(store)

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
