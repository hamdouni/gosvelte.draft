package model

type User struct {
	Username string
	Password string
	Role     Role
}

var Secure Security

// Contrat avec le service de sécurité
type Security interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	HashPassword(pw string) (encryptedPassword string, err error)
	CheckPassword(pw, hashed string) bool
}

var UserStore UserStorage

// Contrat avec le service de stockage
type UserStorage interface {
	GetPasswordUser(username string) (encryptedPassword string)
	AddUser(username, password string, role Role) error
}

func NewUser(un, pw string, ro Role) (*User, error) {
	if len(un) < 4 {
		return nil, ErrUsernameTooShort
	}
	if len(pw) < 4 {
		return nil, ErrPasswordTooShort
	}
	if ro < Customer || ro > Administrator {
		return nil, ErrUndefinedRole
	}

	u := &User{
		Username: un,
		Password: pw,
		Role:     ro,
	}
	return u, nil
}

func CheckPassword(username, password string) bool {
	hashed := UserStore.GetPasswordUser(username)

	return Secure.CheckPassword(password, hashed)
}

func AddUser(username, password string, role Role) error {
	hashed, err := Secure.HashPassword(password)
	if err != nil {
		return err
	}
	UserStore.AddUser(username, hashed, role)

	return nil
}

func Decrypt(message string) (string, error) {
	return Secure.Decrypt(message)
}

func Encrypt(message string) (string, error) {
	return Secure.Encrypt(message)
}
