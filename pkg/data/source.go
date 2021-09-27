package data

const (
	SRC_RANDCHAR string = "randchar"
)

type Source interface {
	HasNext() bool
	GetNext() string
	Size() int
	Reset()
}
