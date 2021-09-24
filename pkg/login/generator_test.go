package login

import (
	"fmt"
	"testing"
	"wpc/pkg/data"
)

func TestLoginGenerator(t *testing.T) {
	gen := Generator{
		usernames: NewSource(makeTestList("user", 13)),
		passwords: NewSource(makeTestList("pass", 12)),
	}
	processed := 0

	conduit := make(chan data.Creds)

	go gen.Generate(conduit)
	batchSize := 0
	for {
		_, ok := <-conduit
		if !ok {
			break
		}
		batchSize++
		processed++
		if batchSize >= 10 {
			batchSize = 0
		}
	}

	if processed != gen.Size() {
		t.Fatalf("expected all to be processed (%d), got %d", gen.Size(), processed)
	}
}

func makeTestList(suffix string, length int) []string {
	res := []string{}
	for i := 1; i <= length; i++ {
		res = append(res, fmt.Sprintf("%s-%d", suffix, i))
	}
	return res
}
