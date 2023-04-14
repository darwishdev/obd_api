package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/darwishdev/obd_api/pkg/util"
)

type sessionCreateTest struct {
	name      string
	session   *SessionCreateParams
	expectErr bool
}

type sessionAttachCodeTest struct {
	name        string
	sessionCode *SessionAttachCodeParams
	expectErr   bool
}

func getValidSession() *SessionCreateParams {
	user := createNewUser()
	return &SessionCreateParams{
		CarID:  user.CarID.Int64,
		IsLive: true,
	}
}

func createNewSession() *Session {
	createdSession, err := testQueries.SessionCreate(context.Background(), *getValidSession())
	if err != nil {
		return nil
	}

	return &createdSession
}

func TestSessionCreate(t *testing.T) {
	// Define a slice of test cases
	testcases := []sessionCreateTest{
		{
			name:      "ValidSession",
			session:   getValidSession(),
			expectErr: false,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the SessionCreate function with the session data from the current test case
			createdSession, err := testQueries.SessionCreate(context.Background(), *tc.session)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			fmt.Println(createdSession)
			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  sessions created during test
			// testQueries.SessionDelete(context.Background(), createdSession.SessionID)

		})
	}
}

func TestSessionAttachCode(t *testing.T) {
	// Define a slice of test cases
	session := createNewSession()
	testcases := []sessionAttachCodeTest{
		{
			name: "ValidSession",
			sessionCode: &SessionAttachCodeParams{
				SessionID: session.SessionID,
				CodeID:    util.RandomInt(1, 20),
			},
			expectErr: false,
		},
		{
			name: "InvalideCodeId",
			sessionCode: &SessionAttachCodeParams{
				SessionID: session.SessionID,
				CodeID:    150,
			},
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the SessionAttachCode function with the session data from the current test case
			createdSession, err := testQueries.SessionAttachCode(context.Background(), *tc.sessionCode)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			fmt.Println(createdSession)
			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  sessions created during test
			// testQueries.SessionDelete(context.Background(), createdSession.SessionID)

		})
	}
}

// type sessionGetTest struct {
// 	name      string
// 	sessionID    int64
// 	expectErr bool
// }

// func TestSessionGet(t *testing.T) {

// 	newSession, err := testQueries.SessionCreate(context.Background(), *getValidSession())
// 	if err != nil {
// 		t.Errorf("Error creating session: %v", err)
// 	}

// 	// Define a slice of test cases
// 	testcases := []sessionGetTest{
// 		{
// 			name:      "ValidSessionID",
// 			sessionID:    newSession.SessionID,
// 			expectErr: false,
// 		},
// 		{
// 			name:      "NonExistentSessionID",
// 			sessionID:    newSession.SessionID + 1,
// 			expectErr: true,
// 		},
// 	}

// 	// Loop through the test cases and test each one
// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			// Call the SessionGet function with the session ID from the current test case
// 			_, err := testQueries.SessionGet(context.Background(), tc.sessionID)

// 			// If the current test case expects an error and no error occurred, fail the test
// 			if tc.expectErr && err == nil {
// 				t.Errorf("Expected an error but got none")
// 			}
// 			// If the current test case does not expect an error and an error occurred, fail the test
// 			if !tc.expectErr && err != nil {
// 				t.Errorf("Expected no error but got %v", err)
// 			}
// 		})
// 	}
// }

// type sessionDeleteTest struct {
// 	name      string
// 	sessionID    int64
// 	expectErr bool
// }

// func TestSessionDelete(t *testing.T) {

// 	newSession, err := testQueries.SessionCreate(context.Background(), *getValidSession())
// 	if err != nil {
// 		t.Errorf("Error creating session: %v", err)
// 	}

// 	// Define a slice of test cases
// 	testcases := []sessionDeleteTest{
// 		{
// 			name:      "ValidSessionID",
// 			sessionID:    newSession.SessionID,
// 			expectErr: false,
// 		},
// 		{
// 			name:      "NonExistentSessionID",
// 			sessionID:    999999,
// 			expectErr: true,
// 		},
// 	}

// 	// Loop through the test cases and test each one
// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			// Call the SessionDelete function with the session ID from the current test case
// 			deletedSession, err := testQueries.SessionDelete(context.Background(), tc.sessionID)

// 			// If the current test case expects an error and no error occurred, fail the test
// 			if tc.expectErr && err == nil {
// 				t.Errorf("Expected an error but got none")
// 			}

// 			// If the current test case does not expect an error and an error occurred, fail the test
// 			if !tc.expectErr && err != nil {
// 				t.Errorf("Expected no error but got %v", err)
// 			}

// 			// If the session was deleted but the deleted_at field is nil, fail the test
// 			if deletedSession.DeletedAt.Time.IsZero() && !tc.expectErr {
// 				t.Errorf("Expected deleted session to have deleted_at field set")
// 			}

// 		})
// 	}
// }
