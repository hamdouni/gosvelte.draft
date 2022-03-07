package model_test

import (
	"admin/model"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name     string
		username string
		password string
		role     model.Role
		err      error
	}{
		{
			name:     "Correct username and password length",
			username: "username",
			password: "password",
			role:     model.Customer,
			err:      nil,
		},
		{
			name:     "Username too short ie less than 4 chars",
			username: "abc",
			password: "12345678",
			err:      model.ErrUsernameTooShort,
		},
		{
			name:     "Password too short ie less than 4 chars",
			username: "abcdefgh",
			password: "123",
			err:      model.ErrPasswordTooShort,
		},
		{
			name:     "Role undefined",
			username: "username",
			password: "password",
			role:     9999,
			err:      model.ErrUndefinedRole,
		},
	}

	for _, test := range testcases {
		_, err := model.NewUser(test.username, test.password, test.role)
		if err != test.err {
			t.Fatalf("Waiting %v but got %v", test.err, err)
		}
	}
}
