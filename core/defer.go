package core

import (
	"fmt"
	"math"
)

// defer is used to envoke a function just before the function returns

func Defer() {
	// a := deferTest()
	// fmt.Println(a)
	fmt.Println(deferTest2())
}

func deferTest() int {
	a := 10
	b := a*10
	defer fmt.Println(b)
	return b*10
}

func deferTest2() bool {
	defer fmt.Println("Function defered")
	a := 10
	if a > 1 {
		num := int(math.Sqrt(float64(a)))
		for i:=2;i<num;i++{
			if a%i==0 {
				return false
			}
		}
		return true
	}
	return false
}