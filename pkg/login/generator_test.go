package login

import (
	"testing"
	"wpc/pkg/data"
)

func TestLoginGenerator(t *testing.T) {
	gen := Generator{
		usernames: NewSource([]string{"user1", "user2", "user3"}),
		passwords: NewSource([]string{"pass1", "pass2", "pass3"}),
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
		if batchSize >= 2 {
			batchSize = 0
		}
	}

	if processed != gen.Size() {
		t.Fatalf("expected all to be processed (%d), got %d", gen.Size(), processed)
	}
}
