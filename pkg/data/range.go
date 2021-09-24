package data

import "sync"

type Range struct {
	sync.Mutex
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
	r.Lock()
	r.pos += 1
	r.Unlock()
}
func (r *Range) Reset() {
	r.Lock()
	r.pos = 0
	r.Unlock()
}

func NewRange(size int) Range {
	return Range{size: size, pos: 0}
}
