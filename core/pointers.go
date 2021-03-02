package core

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

// Pointers in golang
func Pointers() {
	val := 20
	fmt.Println(val)

	ptr := &val

	//print the address where the value of 20 is stored
	fmt.Println(ptr)

	// print the value stored at a address (deferencing)
	fmt.Println(*ptr)

	*ptr = 10
	fmt.Println(*ptr)
	fmt.Println(val)
	c := *ptr
	fmt.Println(c)
	// swap two variables using a swap function
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
