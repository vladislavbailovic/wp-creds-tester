package login

import (
	"wpc/pkg/data"
	"wpc/pkg/login/source"
)

func NewSource(lst []string) data.Source {
	return source.NewFixedList(lst)
}
