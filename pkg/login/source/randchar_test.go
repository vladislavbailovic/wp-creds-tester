package source

import "testing"

func TestRandCharCreation(t *testing.T) {
	src := NewRandChar()

	if src.Size() != RAND_SIZE {
		t.Fatalf("expected default size")
	}

	src2 := NewRandChar(2, 161, 1312)
	if src2.Size() != 2 {
		t.Fatalf("expected size to be 2")
	}
}

func TestRandCharGetNext(t *testing.T) {
	min := 8
	max := 24
	src := NewRandChar(161, min, max)

	for src.HasNext() {
		out := src.GetNext()
		if len(out) < min {
			t.Fatalf("expected out to be longer than %d: [%s] (%d)", min, out, len(out))
		}
		if len(out) > max {
			t.Fatalf("expected out to be longer than %d: [%s] (%d)", max, out, len(out))
		}
	}
}
