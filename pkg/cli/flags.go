package cli

import (
	"flag"
	"os"
	"wpc/pkg/data"
)

func GetOptions(args ...string) *data.Options {
	opts := data.GetOptions()
	cli := flag.NewFlagSet("wp-creds-test", flag.ContinueOnError)
	if len(args) == 0 {
		args = os.Args[1:]
	}

	url := cli.String("url", opts.URL, "Login URL")
	users := cli.String("usr", opts.Usernames, "Usernames source")
	pwds := cli.String("pwd", opts.Passwords, "Passwords source")
	help := cli.Bool("help", false, "Show help")

	cli.Parse(args)
	if *help || *url == "" {
		cli.PrintDefaults()
		return nil
	}

	opts.URL = *url
	opts.Usernames = *users
	opts.Passwords = *pwds

	return opts
}
