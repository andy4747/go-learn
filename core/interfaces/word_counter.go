package interfaces

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error){
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	count := 0
	for input.Scan() {
		count++
	}
	*w = *w + WordCounter(count)
	fmt.Println(count)
	return count,nil
}

func WordCount() {
	var a WordCounter = 0
	a.Write([]byte("Angel Dhakal is my name"))
}
