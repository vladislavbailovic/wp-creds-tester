package main

import "wpc/pkg/data"

func main() {
}

type Source interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}

type FixedListSource struct {
	data.Range
	store []string
}

func (fls *FixedListSource) GetNext() string {
	val := fls.store[fls.Position()]
	fls.Advance()
	return val
}

func NewSource(lst []string) Source {
	return &FixedListSource{
		data.NewRange(len(lst)),
		lst,
	}
}
