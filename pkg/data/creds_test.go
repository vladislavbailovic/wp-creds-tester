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

func TestValidatedCreds(t *testing.T) {
	creds := NewCreds("user", "pass")
	invalid := NewValidatedCreds(creds, false)
	valid := NewValidatedCreds(creds, true)

	if invalid.IsValid() {
		t.Fatalf("expected invalid creds to be invalid")
	}
	if !valid.IsValid() {
		t.Fatalf("expected valid creds to be valid")
	}
}
