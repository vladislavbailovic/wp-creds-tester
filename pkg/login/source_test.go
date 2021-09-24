package login

import "testing"

func TestFactoryReturnsPointer(t *testing.T) {
	test := []string{"one", "two", "three"}
	src := NewSource(test)

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
