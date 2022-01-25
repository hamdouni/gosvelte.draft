package app_test

import (
	"app"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name     string
		username string
		password string
		role     app.RoleType
		err      error
	}{
		{
			name:     "Correct username and password length",
			username: "username",
			password: "password",
			role:     app.Customer,
			err:      nil,
		},
		{
			name:     "Username too short ie less than 4 chars",
			username: "abc",
			password: "12345678",
			err:      app.ErrUsernameTooShort,
		},
		{
			name:     "Password too short ie less than 4 chars",
			username: "abcdefgh",
			password: "123",
			err:      app.ErrPasswordTooShort,
		},
		{
			name:     "Role undefined",
			username: "username",
			password: "password",
			role:     9999,
			err:      app.ErrUndefinedRole,
		},
	}

	for _, test := range testcases {
		_, err := app.NewUser(test.username, test.password, test.role)
		if err != test.err {
			t.Fatalf("Waiting %v but got %v", test.err, err)
		}
	}
}
