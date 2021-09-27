package main

import (
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
		login.NewSource([]string{"user1", "bog", "user3"}),
		login.NewSource([]string{"pass1", "pass2", "bog"}),
	)
	tester := login.NewTester("http://multiwp.test/wp-login.php", generator)

	printer := cli.NewPrintSubscriber()
	tester.Subscribe(login.EVT_START, printer.Header)
	tester.Subscribe(login.EVT_VALIDATED, printer.Item)
	tester.Subscribe(login.EVT_DONE, printer.Footer)

	client := web.NewClient()
	tester.Test(client)
}
