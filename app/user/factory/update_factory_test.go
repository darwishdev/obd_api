package factory

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"

	"github.com/stretchr/testify/require"
)

type newUserUpdateFromProtoTest struct {
	name     string
	req      *obdv1.UserUpdateRequest
	expected *db.UserUpdateParams
	err      error
}
type userUpdateRequestValidationTest struct {
	name     string
	user     *obdv1.UserUpdateRequest
	expected error
}

func getValidUserUpdateRequest(field string, value interface{}) *obdv1.UserUpdateRequest {
	validUser := &obdv1.UserUpdateRequest{
		Name:     util.RandomName(),
		Email:    util.RandomEmail(),
		Phone:    util.RandomPhone(),
		Password: util.RandomString(6),
	}
	if field != "" {
		err := util.SetField(validUser, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}

	return validUser
}

// var ErrInvalidName error
func TestUserUpdateRequestValidation(t *testing.T) {
	tests := []userUpdateRequestValidationTest{
		// Valid user
		{name: "valid", user: getValidUserUpdateRequest("", ""), expected: nil},
		{name: "invalid name", user: getValidUserUpdateRequest("Name", "a"), expected: ErrInvalidName},
		{name: "invalid email too long", user: getValidUserUpdateRequest("Email", util.RandomString(201)), expected: ErrInvalidEmailLength},
		{name: "invalid email too short", user: getValidUserUpdateRequest("Email", util.RandomString(2)), expected: ErrInvalidEmailLength},
		{name: "invalid email - normal string", user: getValidUserUpdateRequest("Email", util.RandomString(6)), expected: ErrInvalidEmail},
		{name: "invalid phone", user: getValidUserUpdateRequest("Phone", util.RandomString(3)), expected: ErrInvalidPhone},
		// {name: "invalid password empty", user: getValidUserUpdateRequest("Password", ""), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserUpdateRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserUpdateRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too short", user: getValidUserUpdateRequest("Password", util.RandomString(5)), expected: ErrInvalidPassword},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := userUpdateRequestValidation(tc.user)
			if tc.expected != nil {
				require.EqualError(t, err, tc.expected.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewUserUpdateFromProto(t *testing.T) {
	validUser := getValidUserUpdateRequest("", "")
	convertedUser := &db.UserUpdateParams{
		Name:  sql.NullString{String: validUser.Name, Valid: true},
		Email: sql.NullString{String: validUser.Email, Valid: true},
		Phone: sql.NullString{String: validUser.Phone, Valid: true},
	}
	tests := []newUserUpdateFromProtoTest{
		// Valid user
		{name: "valid", req: validUser, expected: convertedUser, err: nil},

		// // Invalid name - Too Long
		{name: "Invalid name - Too Long ", req: getValidUserUpdateRequest("Name", util.RandomString(201)), expected: nil, err: ErrInvalidName},
		// // Invalid name - Too Short
		{name: "Invalid name - Too Short ", req: getValidUserUpdateRequest("Name", util.RandomString(1)), expected: nil, err: ErrInvalidName},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("req=%v", tc.req), func(t *testing.T) {
			resp, err := factory.NewUserUpdateFromProto(tc.req)
			if tc.err != nil {
				require.EqualError(t, err, tc.err.Error())
				require.Nil(t, resp)
			} else {
				convertedUser.Password = resp.Password
				require.NoError(t, err)
				require.Equal(t, tc.expected, resp)
			}
		})
	}
}

func TestNewUserUpdateFromSqlResponse(t *testing.T) {

	t.Run("test conversion between db and api responses", func(t *testing.T) {
		resp, err := factory.NewUserCreateFromSqlResponse(validUserResponse)
		require.NoError(t, err)
		require.Equal(t, validUserResponse.Name, resp.User.Name)
		require.Equal(t, validUserResponse.Email, resp.User.Email)
		require.Equal(t, validUserResponse.Phone, resp.User.Phone)
		require.Equal(t, validUserResponse.Password, resp.User.Password)

	})
}
