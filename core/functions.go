package core

import (
	"fmt"
)

// Functions is the demonstration of functions in go
func Functions() {
	// function

	/*
		func functionName(parameter type) reutrn type {
			function body
		}
	*/

	fmt.Println(greet("Angel Dhakal"))

}

func greet(name string) string {
	return "Hello!, " + name
}
