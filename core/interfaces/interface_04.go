package interfaces

import "fmt"

/*
An interface that has zero methods is called an empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.
*/

func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//Interface4 interface
func Interface4() {
	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Angel Dhakal",
	}
	describe2(strt)
}
