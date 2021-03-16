package exercises

import (
	"bufio"
	"fmt"
	"os"
)

func DuplicateLines(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			countLines(f,counts)
			f.Close()
		}
	}
	for line, number := range counts{
		if number > 1 {
			fmt.Printf("%v\t==>\t%v\n",line, number)
		}
	}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	if input.Err() != nil {
		panic(input.Err().Error())
	}
	for input.Scan() {
		counts[input.Text()]++
	}
}
