package data

import "testing"

func TestCreds(t *testing.T) {
	creds := NewCreds("user", "pass")

	if creds.Username() != "user" {
		t.Fatalf("invalid username")
	}
	if creds.Password() != "pass" {
		t.Fatalf("invalid password")
	}
}
