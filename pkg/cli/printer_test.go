package cli

import (
	"testing"
	"wpc/pkg/data"
)

func TestPrinter(t *testing.T) {
	creds := data.NewCreds("user1", "pass1")
	valid := data.NewValidatedCreds(creds, true)
	printer := NewPrinter()

	em := data.NewEmitter()
	em.Subscribe("test", printer.PrintSubscriber)

	em.Publish("test", valid)
}
