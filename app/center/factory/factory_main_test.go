package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	factory CenterFactoryInterface
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewCenterFactory()

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
