package core

import (
	"fmt"
)

//Slice function
func Slice() {
	a := make([]int,0,3)
	a = append(a,10,20,30,40,50)
	fmt.Println(a)
	fmt.Println(cap(a))
}