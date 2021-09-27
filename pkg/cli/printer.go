package cli

import (
	"fmt"
	"wpc/pkg/data"
)

type Printer struct{}

func NewPrinter() Printer {
	return Printer{}
}

func (p Printer) FormatItem(creds data.ValidatedCreds) string {
	valid := "invalid"
	if creds.IsValid() {
		valid = "valid"
	}
	return fmt.Sprintf("%s:%s\t[%s]", creds.Username(), creds.Password(), valid)
}

func (p Printer) PrintItem(creds data.ValidatedCreds) {
	fmt.Println(p.FormatItem(creds))
}

func (p Printer) PrintHeader() {
	fmt.Println("---- start ----")
}

func (p Printer) PrintFooter() {
	fmt.Println("---- end ----")
}

type PrintSubscriber struct {
	Printer
}

func NewPrintSubscriber() PrintSubscriber {
	return PrintSubscriber{NewPrinter()}
}

func (ps PrintSubscriber) Item(evtData []interface{}) {
	var creds data.ValidatedCreds
	for _, item := range evtData {
		creds = item.(data.ValidatedCreds)
		ps.PrintItem(creds)
	}
}

func (ps PrintSubscriber) Items(evtData []interface{}) {
	if len(evtData) > 0 {
		creds := evtData[0].([]data.ValidatedCreds)
		for _, item := range creds {
			ps.PrintItem(item)
		}
	}
}

func (ps PrintSubscriber) Header(evtData []interface{}) {
	ps.PrintHeader()
}

func (ps PrintSubscriber) Footer(evtData []interface{}) {
	ps.PrintFooter()
}
