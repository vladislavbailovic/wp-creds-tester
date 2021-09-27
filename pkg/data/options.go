package data

import "sync"

type Options struct {
	URL       string
	Usernames string
	Passwords string
}

var options *Options
var lock = &sync.Mutex{}

func newOptions() *Options {
	return &Options{}
}

func GetOptions() *Options {
	if options == nil {
		lock.Lock()
		defer lock.Unlock()
		if options == nil {
			options = newOptions()
		}
	}
	return options
}
