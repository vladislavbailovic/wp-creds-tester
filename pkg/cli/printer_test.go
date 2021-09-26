package cli

import (
	"testing"
	"wpc/pkg/data"
)

func TestPrinter(t *testing.T) {
	creds := data.NewCreds("user1", "pass1")
	valid := data.NewValidatedCreds(creds, true)
	printer := NewPrintSubscriber()

	em := data.NewEmitter()
	em.Subscribe("test", printer.Item)
	em.Subscribe("test", printer.Header)
	em.Subscribe("test", printer.Footer)

	em.Publish("test", valid)
}
