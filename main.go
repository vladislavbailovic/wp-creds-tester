package main

func main() {
}

type Range struct {
	size int
	pos  int
}

func (r Range) HasNext() bool {
	return r.pos < r.size
}
func (r Range) Size() int {
	return r.size
}
func (r *Range) Advance() {
	r.pos += 1
}
func (r *Range) Reset() {
	r.pos = 0
}

type Source interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}

type FixedListSource struct {
	Range
	store []string
}

func (fls *FixedListSource) GetNext() string {
	val := fls.store[fls.pos]
	fls.Advance()
	return val
}

func NewSource(data []string) Source {
	return &FixedListSource{
		Range{len(data), 0},
		data,
	}
}
