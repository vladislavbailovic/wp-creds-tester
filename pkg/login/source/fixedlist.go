package source

import (
	"wpc/pkg/data"
)

type FixedList struct {
	data.Range
	store []string
}

func NewFixedList(lst []string) data.Source {
	return &FixedList{
		data.NewRange(len(lst)),
		lst,
	}
}

func (fls *FixedList) GetNext() string {
	val := fls.store[fls.Position()]
	fls.Advance()
	return val
}
