package credential_test

import (
	"testing"
	"wtk/biz"
	"wtk/biz/credential"
	"wtk/ext/secure"
	"wtk/ext/store/ram"
)

func TestCreateCredential(t *testing.T) {
	testcases := []struct {
		realm    string
		name     string
		username string
		password string
		role     credential.Role
		err      error
	}{
		{
			realm:    "FakeRealm",
			name:     "Correct username and password length",
			username: "username",
			password: "password",
			role:     credential.Customer,
			err:      nil,
		},
		{
			realm:    "FakeRealm",
			name:     "Username too short ie less than 4 chars",
			username: "abc",
			password: "12345678",
			err:      credential.ErrUsernameTooShort,
		},
		{
			realm:    "FakeRealm",
			name:     "Password too short ie less than 4 chars",
			username: "abcdefgh",
			password: "123",
			err:      credential.ErrPasswordTooShort,
		},
		{
			realm:    "FakeRealm",
			name:     "Role undefined",
			username: "username",
			password: "password",
			role:     9999,
			err:      credential.ErrUndefinedRole,
		},
	}

	for _, test := range testcases {
		_, err := credential.New(test.realm, test.username, test.password, test.role)
		if err != test.err {
			t.Fatalf("Waiting %v but got %v", test.err, err)
		}
	}
}

func TestDuplicateCredential(t *testing.T) {

	// composant de stockage en RAM
	storage, err := ram.New()
	if err != nil {
		t.Fatalf("Initializing ram: %s", err)
	}

	// composant de sécurité
	security, err := secure.New()
	if err != nil {
		t.Fatalf("Initializing security: %s", err)
	}

	// configure le métier avec le storage et la sécurité
	biz.Initialize(&storage, security)

	// ajoute un user de test
	err = credential.Add("FakeRealm", "test", "test", credential.Administrator)
	if err != nil {
		t.Fatalf("Creating test user: %s", err)
	}

	// ajoute le même user
	err = credential.Add("FakeRealm", "test", "test", credential.Administrator)
	if err != credential.ErrUsernameUsed {
		t.Fatalf("waiting %s got %s", credential.ErrUsernameUsed, err)
	}
}
