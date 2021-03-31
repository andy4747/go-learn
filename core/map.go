package core

import (
	"bufio"
	"fmt"
	"os"
)

// Maps in golang
func Maps() {
	data := make(map[string]int)
	data["a"] = 0
	data["b"] = 1
	fmt.Println(data)

	// print the length of the map
	fmt.Println(len(data))

	// delete a key/value pair from map
	delete(data, "a")
	fmt.Println(data)

	// access the absent key
	_, isPresent := data["c"]
	fmt.Println(isPresent)
}

func equal(x,y map[string]int)bool{
	if len(x) != len(y){
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}


func Dedup(){
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	}
}
