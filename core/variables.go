package core

import (
	"fmt"
)

// Variables is a function that uses two type of variable initialization
func Variables() {
	// implicit type for a variable
	var x int = 10
	var y int = 20
	fmt.Println(x + y)

	// explicit type for a variable
	a := 10
	b := 20
	fmt.Println(a + b)
}
