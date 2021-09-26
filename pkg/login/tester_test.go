package login

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wpc/pkg/data"
	"wpc/pkg/web"
)

func TestTesterDoesNotValidateNonRedirects(t *testing.T) {
	g := &Generator{
		usernames: NewSource([]string{"user1", "user2", "user3"}),
		passwords: NewSource([]string{"pass1", "pass2", "pass3"}),
	}
	c := web.NewClient()

	runTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			sendErrorHeader(w, "error parsing form")
			return
		}

		username := r.Form.Get("log")
		if !strings.HasPrefix(username, "user") {
			sendErrorHeader(w, "expected username, got "+username)
			return
		}

		password := r.Form.Get("pwd")
		if !strings.HasPrefix(password, "pass") {
			sendErrorHeader(w, "expected password, got "+password)
			return
		}

		w.Header().Set("all-good", username+"/"+password)
	}, func(server *httptest.Server) {
		test := NewTester(server.URL, g)
		result := test.Test(c)
		for _, actual := range result {
			if actual.IsValid() {
				t.Fatalf("all combos should be invalid: %s/%s", actual.Username(), actual.Password())
			}
		}
	})

}

func TestTesterValidatesAdminRedirects(t *testing.T) {
	g := &Generator{
		usernames: NewSource([]string{"invalid1", "valid", "invalid2"}),
		passwords: NewSource([]string{"badpass1", "goodpass", "badpass2"}),
	}
	c := web.NewClient()

	runTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			sendErrorHeader(w, "error parsing form")
			return
		}

		username := r.Form.Get("log")
		password := r.Form.Get("pwd")

		if username == "valid" && password == "goodpass" {
			w.Header().Set("location", "/wp-admin/")
			w.WriteHeader(302)
			return
		}

		w.Header().Set("all-good", username+"/"+password)
	}, func(server *httptest.Server) {
		test := NewTester(server.URL, g)
		result := test.Test(c)
		for _, actual := range result {
			if !actual.IsValid() {
				continue
			}
			if actual.Username() != "valid" {
				t.Fatalf("expected username valid, got %s", actual.Username())
			}
			if actual.Password() != "goodpass" {
				t.Fatalf("expected pass goodpass, got %s", actual.Password())
			}
		}
	})
}

func TestTesterDoesNotValidateNonAdminRedirects(t *testing.T) {
	g := &Generator{
		usernames: NewSource([]string{"user1", "user2", "user3"}),
		passwords: NewSource([]string{"pass1", "pass2", "pass3"}),
	}
	c := web.NewClient()

	runTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("location", "/whatever/")
		w.WriteHeader(302)
		return
	}, func(server *httptest.Server) {
		test := NewTester(server.URL, g)
		result := test.Test(c)
		for _, actual := range result {
			if actual.IsValid() {
				t.Fatalf("all combos should be invalid: %s/%s", actual.Username(), actual.Password())
			}
		}
	})
}

func TestTesterEmitsValidatedEvent(t *testing.T) {
	g := &Generator{
		usernames: NewSource([]string{"user1", "user2", "user3"}),
		passwords: NewSource([]string{"pass1", "pass2", "pass3"}),
	}
	c := web.NewClient()
	handler := func(evtData []interface{}) {
		actual := evtData[0].(data.ValidatedCreds)
		if actual.IsValid() {
			t.Fatalf("expected invalid creds: %s/%s", actual.Username(), actual.Password())
		}
	}

	runTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("location", "/whatever/")
		w.WriteHeader(302)
		return
	}, func(server *httptest.Server) {
		test := NewTester(server.URL, g)
		test.Subscribe(EVT_VALIDATED, handler)
		test.Test(c)
	})
}

func sendErrorHeader(w http.ResponseWriter, msg string) {
	w.Header().Set("error", msg)
	w.WriteHeader(500)
	w.Write([]byte(msg))
}

func runTestServer(serverHandler http.HandlerFunc, clientHandler func(*httptest.Server)) {
	server := httptest.NewServer(serverHandler)
	defer server.Close()

	clientHandler(server)
}
