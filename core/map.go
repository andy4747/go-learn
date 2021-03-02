package core

import (
	"fmt"
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
