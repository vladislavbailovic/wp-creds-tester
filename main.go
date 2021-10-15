package main

import (
	"fmt"
	"os"
	"wpc/pkg/cli"
	"wpc/pkg/login"
	"wpc/pkg/web"
)

func main() {
	opts := cli.GetOptions()
	if nil == opts {
		os.Exit(1)
	}
	generator := login.NewGenerator(
		login.NewSource(opts.Usernames),
		login.NewSource(opts.Passwords),
	)
	tester := login.NewTester(opts.URL, generator)

	printer := cli.NewPrintSubscriber()
	tester.Subscribe(login.EVT_START, printer.Header)
	// tester.Subscribe(login.EVT_VALIDATED, printer.Item) // Follow:true
	tester.Subscribe(login.EVT_DONE, printer.Items) // Follow:false
	tester.Subscribe(login.EVT_DONE, printer.Footer)

	client := web.NewClient()
	tester.Test(client)

	fmt.Println(tester.GetStart())
	fmt.Println(tester.GetEnd())
	fmt.Println(tester.GetDuration().Milliseconds())
}
