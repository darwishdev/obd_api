package factory

import (
	"fmt"
	"log"
	"testing"

	obdv1 "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	"github.com/darwishdev/obd_api/pkg/util"

	"github.com/stretchr/testify/require"
)

type userLoginRequestValidationTest struct {
	name     string
	user     *obdv1.UserLoginRequest
	expected error
}

func getValidUserLoginRequest(field string, value interface{}) *obdv1.UserLoginRequest {
	validUser := &obdv1.UserLoginRequest{
		Email:    util.RandomEmail(),
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

func TestUserLoginRequestValidation(t *testing.T) {
	tests := []userLoginRequestValidationTest{
		// Valid user
		{name: "valid", user: getValidUserLoginRequest("", ""), expected: nil},
		{name: "invalid email too long", user: getValidUserLoginRequest("Email", util.RandomString(201)), expected: ErrInvalidEmailLength},
		{name: "invalid email too short", user: getValidUserLoginRequest("Email", util.RandomString(2)), expected: ErrInvalidEmailLength},
		{name: "invalid email - normal string", user: getValidUserLoginRequest("Email", util.RandomString(6)), expected: ErrInvalidEmail},
		{name: "invalid password empty", user: getValidUserLoginRequest("Password", ""), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserLoginRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too long", user: getValidUserLoginRequest("Password", util.RandomString(201)), expected: ErrInvalidPassword},
		{name: "invalid password too short", user: getValidUserLoginRequest("Password", util.RandomString(5)), expected: ErrInvalidPassword},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("name=%v", tc.name), func(t *testing.T) {

			err := factory.UserLoginRequestValidation(tc.user)
			if tc.expected != nil {
				require.EqualError(t, err, tc.expected.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
