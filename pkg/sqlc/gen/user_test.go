package db

import (
	"context"
	"testing"

	"github.com/darwish/obd_api/pkg/util"
)

type userCreateTest struct {
	name      string
	user      *UserCreateParams
	expectErr bool
}

func getValidUser() *UserCreateParams {
	return &UserCreateParams{
		Name:     util.RandomName(),
		Phone:    util.RandomPhone(),
		Email:    util.RandomEmail(),
		Password: util.RandomString(10),
	}
}
func TestUserCreate(t *testing.T) {
	// Define a slice of test cases
	testcases := []userCreateTest{
		{
			name:      "ValidUser",
			user:      getValidUser(),
			expectErr: false,
		},
		{
			name: "TooLongUserName",
			user: &UserCreateParams{
				Name:     util.RandomString(201),
				Phone:    util.RandomPhone(),
				Email:    util.RandomEmail(),
				Password: util.RandomString(10),
			},
			expectErr: true,
		},
		{
			name: "TooLongPhone",
			user: &UserCreateParams{
				Name:     util.RandomName(),
				Phone:    util.RandomString(201),
				Email:    util.RandomEmail(),
				Password: util.RandomString(10),
			},
			expectErr: true,
		},
		{
			name: "TooLongEmail",
			user: &UserCreateParams{
				Name:     util.RandomName(),
				Phone:    util.RandomPhone(),
				Email:    util.RandomString(201),
				Password: util.RandomString(10),
			},
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the UserCreate function with the user data from the current test case
			createdUser, err := testQueries.UserCreate(context.Background(), *tc.user)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}
			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  users created during test
			testQueries.UserDelete(context.Background(), createdUser.UserID)

		})
	}
}

type userGetTest struct {
	name      string
	userID    int64
	expectErr bool
}

func TestUserGet(t *testing.T) {

	newUser, err := testQueries.UserCreate(context.Background(), *getValidUser())
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Define a slice of test cases
	testcases := []userGetTest{
		{
			name:      "ValidUserID",
			userID:    newUser.UserID,
			expectErr: false,
		},
		{
			name:      "NonExistentUserID",
			userID:    newUser.UserID + 1,
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the UserGet function with the user ID from the current test case
			_, err := testQueries.UserGet(context.Background(), tc.userID)

			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}
			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}

type userDeleteTest struct {
	name      string
	userID    int64
	expectErr bool
}

func TestUserDelete(t *testing.T) {

	newUser, err := testQueries.UserCreate(context.Background(), *getValidUser())
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Define a slice of test cases
	testcases := []userDeleteTest{
		{
			name:      "ValidUserID",
			userID:    newUser.UserID,
			expectErr: false,
		},
		{
			name:      "NonExistentUserID",
			userID:    999999,
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the UserDelete function with the user ID from the current test case
			deletedUser, err := testQueries.UserDelete(context.Background(), tc.userID)

			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}

			// If the user was deleted but the deleted_at field is nil, fail the test
			if deletedUser.DeletedAt.Time.IsZero() && !tc.expectErr {
				t.Errorf("Expected deleted user to have deleted_at field set")
			}

		})
	}
}
