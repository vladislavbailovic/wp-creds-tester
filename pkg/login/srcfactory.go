package login

import (
	"strconv"
	"strings"
	"wpc/pkg/data"
	"wpc/pkg/login/source"
)

func NewSource(lst ...string) data.Source {
	if len(lst) == 1 {
		return newSingleSource(lst[0])
	}
	return source.NewFixedList(lst)
}

func newSingleSource(src string) data.Source {
	if strings.HasPrefix(src, "gen:") {
		return newGeneratorSource(src)
	}
	// ...
	// We exhausted other options, let's assume that's all there is.
	return source.NewFixedList([]string{src})
}

func newGeneratorSource(src string) data.Source {
	parts := strings.Split(src, ":")
	if len(parts) < 2 {
		return source.NewFixedList([]string{src})
	}

	if parts[1] == data.SRC_RANDCHAR {
		args := []int{}
		if len(parts) > 2 {
			size, err := strconv.Atoi(parts[2])
			if err == nil {
				args = append(args, size)
			}
		}
		if len(parts) > 3 {
			minLen, err := strconv.Atoi(parts[3])
			if err == nil {
				args = append(args, minLen)
			}
		}
		if len(parts) > 4 {
			maxLen, err := strconv.Atoi(parts[4])
			if err == nil {
				args = append(args, maxLen)
			}
		}
		return source.NewRandChar(args...)
	}

	return source.NewFixedList(parts)
}
