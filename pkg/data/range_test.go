package data

import "testing"

func TestRange(t *testing.T) {
	r := NewRange(2)

	if 2 != r.Size() {
		t.Fatalf("size should be 2")
	}
	if !r.HasNext() {
		t.Fatalf("should have next initially")
	}

	r.Advance()
	if !r.HasNext() {
		t.Fatalf("should have next at first")
	}
	if r.Position() != 1 {
		t.Fatalf("expected position to be 1, got %d", r.Position())
	}

	r.Advance()
	if r.HasNext() {
		t.Fatalf("should be exhausted")
	}
	if r.Position() != 2 {
		t.Fatalf("expected position to be 2, got %d", r.Position())
	}

	r.Reset()
	if !r.HasNext() {
		t.Fatalf("should have next after reset")
	}
}
