package data

type Options struct {
	URL       string
	Usernames string
	Passwords string
}

func getDefaultOptions() *Options {
	return &Options{}
}

func GetOptions() *Options {
	return getDefaultOptions()
}
