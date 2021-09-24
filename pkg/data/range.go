package data

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
func (r Range) Position() int {
	return r.pos
}
func (r *Range) Advance() {
	r.pos += 1
}
func (r *Range) Reset() {
	r.pos = 0
}

func NewRange(size int) Range {
	return Range{size, 0}
}
