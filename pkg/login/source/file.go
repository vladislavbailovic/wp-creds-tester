package source

import (
	"bufio"
	"os"
	"strings"
	"wpc/pkg/data"
)

type File struct {
	FixedList
}

func NewFile(filename string) data.Source {
	fp, err := os.Open(filename)
	if err != nil {
		return &File{*NewFixedList([]string{filename}).(*FixedList)}
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	lines := []string{}
	for {
		line, err := reader.ReadString([]byte("\n")[0])
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	list := NewFixedList(lines)
	return &File{*list.(*FixedList)}
}
