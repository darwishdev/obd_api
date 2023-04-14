package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	factory            SessionFactoryInterface
	ErrInvalidCenterId error
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewSessionFactory()

	// Cleanup, if necessary

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
