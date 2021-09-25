package login

import (
	"net/http"
	"strings"
	"sync"
	"wpc/pkg/data"
	"wpc/pkg/web"
)

const (
	DEFAULT_THREADS int = 5
)

type Tester struct {
	url       string
	generator *Generator
	threads   int
}

func NewTester(url string, generator *Generator) Tester {
	return Tester{url, generator, DEFAULT_THREADS}
}

func (t *Tester) SetThreads(threads int) {
	t.threads = threads
}

func (t Tester) Test(client *web.Client) []data.ValidatedCreds {
	conduit := make(chan data.Creds)
	result := []data.ValidatedCreds{}
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
			result = append(result, res)
			// fmt.Println(res)
		}(creds)
		pool.Advance()
		if !pool.HasNext() {
			pool.Reset()
			// fmt.Printf("----- wait %d -----\n", pool.Position())
			wg.Wait()
		}
	}
	wg.Wait()

	return result
}

func (t Tester) validateCreds(creds data.Creds, client *web.Client) data.ValidatedCreds {
	response := client.Request(t.url, creds)
	if isLoginSuccessful(response) {
		return data.NewValidatedCreds(creds, true)
	}

	return data.NewValidatedCreds(creds, false)
}

func isRedirect(response *http.Response) bool {
	return response.StatusCode > 300 && response.StatusCode < 400
}

func isRedirectingToAdmin(response *http.Response) bool {
	if !isRedirect(response) {
		return false
	}

	location, err := response.Location()
	if err != nil {
		return false
	}
	if !strings.Contains(location.String(), "wp-admin") {
		return false
	}

	return true
}

func hasLoggedInCookie(response *http.Response) bool {
	hasCookie := false
	for _, cookie := range response.Cookies() {
		if strings.HasPrefix(cookie.Name, "wordpress_logged_in") {
			hasCookie = true
		}
	}
	return hasCookie
}

func isLoginSuccessful(response *http.Response) bool {
	return isRedirectingToAdmin(response) && hasLoggedInCookie(response)
}
