package data

type Creds struct {
	username string
	password string
}

func (l Creds) Username() string {
	return l.username
}
func (l Creds) Password() string {
	return l.password
}

func NewCreds(username, password string) Creds {
	return Creds{username, password}
}

type ValidatedCreds struct {
	Creds
	valid bool
}

func (vc ValidatedCreds) IsValid() bool {
	return vc.valid
}

func NewValidatedCreds(creds Creds, valid bool) ValidatedCreds {
	return ValidatedCreds{creds, valid}
}
