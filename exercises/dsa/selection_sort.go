package dsa

import (
	"math/rand"
	"time"
)

//GenerateSlice regerates a slice of int of given size
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//SelectionSort sorts unordered slice in time complexity of o(n^2)
func SelectionSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}
