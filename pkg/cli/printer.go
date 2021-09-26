package cli

import (
	"fmt"
	"wpc/pkg/data"
)

type Printer struct{}

func NewPrinter() Printer {
	return Printer{}
}

func (p Printer) Format(creds data.ValidatedCreds) string {
	valid := "invalid"
	if creds.IsValid() {
		valid = "valid"
	}
	return fmt.Sprintf("%s:%s\t[%s]", creds.Username(), creds.Password(), valid)
}

func (p Printer) Print(creds data.ValidatedCreds) {
	fmt.Println(p.Format(creds))
}

func (p Printer) PrintSubscriber(eventData []interface{}) {
	var creds data.ValidatedCreds
	for _, item := range eventData {
		creds = item.(data.ValidatedCreds)
		p.Print(creds)
	}
}
