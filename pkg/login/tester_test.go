package login

import (
	"testing"
	"wpc/pkg/web"
)

func TestTesterBatching(t *testing.T) {
	g := &Generator{
		usernames: NewSource([]string{"user1", "user2", "user3"}),
		passwords: NewSource([]string{"pass1", "pass2", "pass3"}),
	}
	c := web.NewClient()
	test := NewTester(g)

	test.Test(c)
}
