package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	factory WinchFactoryInterface
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here
	factory = NewWinchFactory()

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
