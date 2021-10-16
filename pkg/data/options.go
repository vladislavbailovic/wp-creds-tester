package data

type Options struct {
	URL       string
	Usernames string
	Passwords string
	Follow    bool
}

func getDefaultOptions() *Options {
	return &Options{}
}

func GetOptions() *Options {
	return getDefaultOptions()
}
