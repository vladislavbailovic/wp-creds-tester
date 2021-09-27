package cli

import (
	"testing"
)

func TestFlagsAreNilIfNoUrl(t *testing.T) {
	opts := GetOptions()
	if opts != nil {
		t.Fatalf("expected opts to be nil if no URL set")
	}

	opts2 := GetOptions("--help", "-url", "test")
	if opts2 != nil {
		t.Fatalf("expected opts to be nil when help is asked for")
	}
}

func TestFlagsAreSetFromCliArgs(t *testing.T) {
	args := []string{
		"-url", "test.com",
	}
	opts := GetOptions(args...)
	if opts.URL != "test.com" {
		t.Fatalf("URL should be set")
	}
	if opts.Usernames != "" {
		t.Fatalf("Usernames should be empty at first")
	}
	if opts.Passwords != "" {
		t.Fatalf("Passwords should be empty at first")
	}
}
