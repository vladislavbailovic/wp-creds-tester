package login

import (
	"testing"
	"wpc/pkg/login/source"
)

func TestFactoryReturnsPointer(t *testing.T) {
	test := []string{"one", "two", "three"}
	src := NewSource(test...)

	if 3 != src.Size() {
		t.Fatalf("size should be 3")
	}
	if !src.HasNext() {
		t.Fatalf("should have next initially")
	}

	for idx, expected := range test {
		actual := src.GetNext()
		if actual != expected {
			t.Fatalf("expected %s, but got %s at %d", expected, actual, idx)
		}
		if idx < len(test)-1 && !src.HasNext() {
			t.Fatalf("should have next at pos %d", idx)
		}
	}

	if src.HasNext() {
		t.Fatalf("should be exhausted after iteration")
	}

	src.Reset()
	if !src.HasNext() {
		t.Fatalf("should have next after reset")
	}
}

func TestGetGeneratorSourceFallback(t *testing.T) {
	src := NewSource("gen")
	if src.Size() != 1 {
		t.Fatalf("expected size fail: %d", src.Size())
	}

	src2 := NewSource("gen:unknown:whatever:this:is")
	if src2.Size() != 5 {
		t.Fatalf("expected size fail: %d", src2.Size())
	}
}

func TestGetRandCharGeneratorSource(t *testing.T) {
	src := NewSource("gen:randchar")
	if src.Size() != source.RAND_SIZE {
		t.Fatalf("expected size to be %d, got %d", source.RAND_SIZE, src.Size())
	}

	src2 := NewSource("gen:randchar:161:6:24")
	if src2.Size() != 161 {
		t.Fatalf("expected size to be 161, got %d", src2.Size())
	}
	for src2.HasNext() {
		out := src2.GetNext()
		if len(out) < 6 {
			t.Fatalf("expected output to be longer than 6: [%s](%d)", out, len(out))
		}
		if len(out) > 24 {
			t.Fatalf("expected output to be shorter than 24: [%s](%d)", out, len(out))
		}
	}
}
