package factory

import (
	"fmt"
	"log"
	"testing"
	"time"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	db "github.com/darwishdev/obd_api/pkg/sqlc/gen"
	"github.com/darwishdev/obd_api/pkg/util"

	"github.com/stretchr/testify/require"
)

type newUserFromProtoTest struct {
	name     string
	req      *obdv1.UserCreateRequest
	expected *db.UserCreateParams
	err      error
}
type userCreateRequestValidationTest struct {
	name     string
	user     *obdv1.UserCreateRequest
	expected error
}

func getValidUserRequest(field string, value interface{}) *obdv1.UserCreateRequest {
	validUser := &obdv1.UserCreateRequest{
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

func TestUserCreateRequestValidation(t *testing.T) {
	tests := []userCreateRequestValidationTest{
		// Valid user
		{name: "valid", user: getValidUserRequest("", ""), expected: nil},
		{name: "invalid name", user: getValidUserRequest("Name", "a"), expected: ErrInvalidName},
		{name: "invalid email too long", user: getValidUserRequest("Email", util.RandomString(201)), expected: ErrInvalidEmailLength},
		{name: "invalid email too short", user: getValidUserRequest("Email", util.RandomString(2)), expected: ErrInvalidEmailLength},
		{name: "invalid email - normal string", user: getValidUserRequest("Email", util.RandomString(6)), expected: ErrInvalidEmail},
		{name: "invalid phone", user: getValidUserRequest("Phone", util.RandomString(3)), expected: ErrInvalidPhone},
		{name: "invalid password empty", user: getValidUserRequest("Password", ""), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too short", user: getValidUserRequest("Password", util.RandomString(5)), expected: ErrInvalidPassword},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := userCreateRequestValidation(tc.user)
			if tc.expected != nil {
				require.EqualError(t, err, tc.expected.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewUserFromProto(t *testing.T) {
	validUser := getValidUserRequest("", "")
	convertedUser := &db.UserCreateParams{
		Name:  validUser.Name,
		Email: validUser.Email,
		Phone: validUser.Phone,
	}
	tests := []newUserFromProtoTest{
		// Valid user
		{name: "valid", req: validUser, expected: convertedUser, err: nil},
		// // Invalid name - empty string
		{name: "empty name ", req: getValidUserRequest("Name", ""), expected: nil, err: ErrInvalidName},
		// // Invalid name - Too Long
		{name: "Invalid name - Too Long ", req: getValidUserRequest("Name", util.RandomString(201)), expected: nil, err: ErrInvalidName},
		// // Invalid name - Too Short
		{name: "Invalid name - Too Short ", req: getValidUserRequest("Name", util.RandomString(1)), expected: nil, err: ErrInvalidName},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("req=%v", tc.req), func(t *testing.T) {
			resp, err := factory.NewUserFromProto(tc.req)
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

func TestNewUserFromSqlResponse(t *testing.T) {
	validUserResponse := &db.User{
		Name:      util.RandomName(),
		Email:     util.RandomEmail(),
		Phone:     util.RandomPhone(),
		Password:  util.RandomString(6),
		CreatedAt: time.Now(),
	}

	t.Run("test conversion between db and api responses", func(t *testing.T) {
		resp, err := factory.NewUserCreateFromSqlResponse(validUserResponse)
		require.NoError(t, err)
		require.Equal(t, validUserResponse.Name, resp.User.Name)
		require.Equal(t, validUserResponse.Email, resp.User.Email)
		require.Equal(t, validUserResponse.Phone, resp.User.Phone)
		require.Equal(t, validUserResponse.Password, resp.User.Password)

	})
}
