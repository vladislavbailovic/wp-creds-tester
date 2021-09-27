package data

type Source interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}
