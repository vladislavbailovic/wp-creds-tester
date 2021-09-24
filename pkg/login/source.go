package login

import (
	"wpc/pkg/data"
)

type Source interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}

type FixedList struct {
	data.Range
	store []string
}

func (fls *FixedList) GetNext() string {
	val := fls.store[fls.Position()]
	fls.Advance()
	return val
}

func NewSource(lst []string) Source {
	return &FixedList{
		data.NewRange(len(lst)),
		lst,
	}
}
