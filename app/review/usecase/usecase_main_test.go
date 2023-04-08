package usecase

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	// Initialize the factory instance here

	// Run the tests
	code := m.Run()
	// Exit with the test result code
	os.Exit(code)
}
