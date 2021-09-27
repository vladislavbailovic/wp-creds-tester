package source

import (
	"math/rand"
	"time"
	"wpc/pkg/data"
)

const (
	RAND_SIZE   int    = 10
	RAND_MINLEN int    = 4
	RAND_MAXLEN int    = 12
	RAND_CHARS  string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type RandChar struct {
	data.Range
	minLen int
	maxLen int
}

func NewRandChar(args ...int) data.Source {
	size := RAND_SIZE
	if len(args) > 0 {
		size = args[0]
	}

	minLen := RAND_MINLEN
	if len(args) > 1 {
		minLen = args[1]
	}

	maxLen := RAND_MAXLEN
	if len(args) > 2 {
		maxLen = args[2]
	}

	return &RandChar{
		data.NewRange(size),
		minLen,
		maxLen,
	}
}

func (rc *RandChar) GetNext() string {
	val := rc.GenerateRandomString()
	rc.Advance()
	return val
}

func (rc *RandChar) GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(rc.maxLen-rc.minLen) + rc.minLen
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(RAND_CHARS))
		b[i] = RAND_CHARS[idx]
	}
	return string(b)
}
