package login

import (
	"fmt"
	"sync"
	"wpc/pkg/data"
	"wpc/pkg/web"
)

type Tester struct {
	generator *Generator
	threads   int
}

func NewTester(generator *Generator) Tester {
	return Tester{generator, 2}
}

func (t Tester) Test(client *web.Client) {
	conduit := make(chan data.Creds)
	var wg sync.WaitGroup

	go t.generator.Generate(conduit)
	pool := data.NewRange(t.threads)
	for {
		creds, ok := <-conduit
		if !ok {
			break
		}

		wg.Add(1)
		go func(c data.Creds) {
			defer wg.Done()
			res := t.validateCreds(c, client)
			fmt.Println(res)
		}(creds)
		pool.Advance()
		if !pool.HasNext() {
			pool.Reset()
			fmt.Printf("----- wait %d -----\n", pool.Position())
			wg.Wait()
		}
	}
	wg.Wait()
}

func (t Tester) validateCreds(creds data.Creds, client *web.Client) string {
	return client.Request(creds)
}
