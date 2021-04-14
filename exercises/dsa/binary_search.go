package dsa

import (
	"fmt"
)

func binarySearch(sortedList []int, item int) int {
	firstIndex := 0
	lastIndex := len(sortedList) - 1
	midIndex := int((firstIndex + lastIndex) / 2)
	return midIndex
}

//BinarySearch function
func BinarySearch() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80}
	b := binarySearch(a, 60)
	fmt.Println(b)
}
