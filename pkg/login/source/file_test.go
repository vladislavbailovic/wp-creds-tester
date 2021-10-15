package source

import (
	"fmt"
	"testing"
)

func TestFileCreation(t *testing.T) {
	test := NewFile("../../../fixtures/10lines.txt")

	if test.Size() != 10 {
		t.Fatalf("expected size to be 10")
	}

	for i := 1; i < 10; i++ {
		expected := fmt.Sprintf("Line number %d", i)
		actual := test.GetNext()
		if expected != actual {
			t.Fatalf("line mismatch: expected [%s] but got [%s]", expected, actual)
		}
	}
}
