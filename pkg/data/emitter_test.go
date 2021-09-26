package data

import (
	"testing"
)

func TestEmitterCasting(t *testing.T) {
	var (
		em      Emitter
		source  ValidatedCreds
		handler Subscriber
	)
	for _, state := range []bool{true, false} {
		em = NewEmitter()
		source = NewValidatedCreds(NewCreds("invalid", "invalid"), state)
		handler = func(data []interface{}) {
			creds := data[0].(ValidatedCreds)
			expected := "valid"
			if !state {
				expected = "invalid"
			}
			if state != creds.IsValid() {
				t.Fatalf("expected creds to be %s: %s/%s", expected, creds.Username(), creds.Password())
			}
		}
		em.Subscribe(EventType("test"), handler)
		em.Publish("test", source)
	}
}

func TestSubscriberStorage(t *testing.T) {
	store := NewSubscriberStorage()
	evt := func(data []interface{}) {
		return
	}
	key := EventType("test key")

	if len(store.Get(key)) != 0 {
		t.Fatalf("expected store length to initially be 0")
	}

	store.Add(key, evt)
	res := store.Get(key)
	if len(res) != 1 {
		t.Fatalf("expected 1 event after add")
	}
}
