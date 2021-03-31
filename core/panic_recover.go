package core

import (
	"fmt"
)

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ",r)
	}
	fmt.Println("Hello World")
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: firstName cannot be null")
	} else if lastName == nil {
		panic("runtime error: lastName cannot be null")
	}
	fmt.Println(*firstName, *lastName)
}

func PanicRecover() {
	firstName := "Elon"
	// lastName := "Musk"
	fullName(&firstName, nil)
}