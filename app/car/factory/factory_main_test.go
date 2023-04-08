package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	factory            CarFactoryInterface
	ErrInvalidCenterId error
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewCarFactory()

	// Cleanup, if necessary

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
