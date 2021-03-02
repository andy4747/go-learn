package core

import (
	"errors"
	"fmt"
)

// Errors function
func Errors() {
	fmt.Println("Error handling in golang")
	a, err := sumOf(10, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func sumOf(start, end int) (int, error) {
	total := 0
	if start > end {
		return total, errors.New("end is greater than start")
	}
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
