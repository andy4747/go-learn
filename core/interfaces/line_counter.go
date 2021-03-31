package interfaces

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanLines)
	count := 0
	for input.Scan() {
		count++
	}
	*l += LineCounter(count)
	fmt.Println(count)
	return count, nil
}

func LineCount() {
	var a LineCounter = 0
	a.Write([]byte("Hello\nWorld\nthis\nis\nNepal"))
}
