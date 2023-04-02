package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var factory UserFactoryInterface

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewUserFactory()

	// Run the tests
	code := m.Run()

	// Cleanup, if necessary

	// Exit with the test result code
	os.Exit(code)
}
