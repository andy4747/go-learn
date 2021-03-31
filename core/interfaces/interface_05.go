package interfaces

import "fmt"

//Interface5 function
func Interface5() {
	var a interface{} = 'A'
	// var b interface{} = "Hello"
	fmt.Printf("%T\n", a)
	assert(a)
	// assert(b)
}

func assert(i interface{}) {
	v, ok := i.(int) // get the underlying int value of i
	fmt.Println(v, ok)
}
