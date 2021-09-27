package data

import "testing"

func TestOptionsIsASingleton(t *testing.T) {
	opts := GetOptions()
	opts.URL = "test"

	opts2 := GetOptions()
	if opts.URL != opts2.URL {
		t.Fatalf("expected options to match")
	}
}
