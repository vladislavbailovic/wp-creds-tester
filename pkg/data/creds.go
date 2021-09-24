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
