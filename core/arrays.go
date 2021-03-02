package core

import (
	"fmt"
)

// Arrays function is the demonstration of arrays in golang
func Arrays() {
	// An array variable denotes the entire array; unlike C, array name is not a pointer to the first array element.
	// This means that when we assign or pass around an array value we will make a copy of its contents.
	// To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.

	//example

	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Mango"

	fmt.Println(fruits)

	// declaring and assigning at the same time

	otherFruits := [2]string{"Apple", "Mango"}
	fmt.Println(otherFruits)

	// using int array
	ages := [2]int{19, 19}
	fmt.Println(ages)

	// 2d array
	var nums [5][5]int
	count := 10
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			nums[i][j] = count
			count = count + 1
		}
	}
	fmt.Println(nums)

	// for i := 0; i < len(nums); i++ {
	// 	for j := 0; j < len(nums[i]); j++ {
	// 		fmt.Println(nums[i][j])
	// 	}
	// 	fmt.Println("Iter Complete")
	// }

	// using range to print the array
	for _, row := range nums {
		for _, val := range row {
			fmt.Println(val)
		}
		fmt.Println("Iter Complete")
	}
}
