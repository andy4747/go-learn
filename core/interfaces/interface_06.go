package interfaces

import "fmt"

/*
A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.
*/

//Interface6 function
func Interface6() {
	findType("test")
	findType(50)
	findType(69.69)
}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("Integer type %v\n", i.(int))
	case string:
		fmt.Printf("String type %v\n", i.(string))
	default:
		fmt.Printf("Unknown Type\n")
	}
}
