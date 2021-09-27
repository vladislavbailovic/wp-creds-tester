package login

import (
	"wpc/pkg/data"
)

type Generator struct {
	usernames data.Source
	passwords data.Source
}

func NewGenerator(usernames, passwords data.Source) *Generator {
	return &Generator{usernames, passwords}
}

func (gen *Generator) Generate(conduit chan data.Creds) {
	defer close(conduit)

	for gen.usernames.HasNext() {
		user := gen.usernames.GetNext()
		for gen.passwords.HasNext() {
			pass := gen.passwords.GetNext()
			conduit <- data.NewCreds(user, pass)
		}
		gen.passwords.Reset()
	}
}

func (gen *Generator) Size() int {
	return gen.usernames.Size() * gen.passwords.Size()
}
