package interfaces

import (
	"fmt"
	"strings"
)

/*
In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface. It is much similar to the OOP world. Interface specifies what methods a type should have and the type decides how to implement these methods.

For example WashingMachine can be an interface with method signatures Cleaning() and Drying(). Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.
*/

//Interface in go
func Interface() {
	name := MyString("Angel Dhakal")
	var v VolwelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}

//VolwelsFinder interface definition
type VolwelsFinder interface {
	FindVowels() []rune
}

// MyString is a type alias for string
type MyString string

func (ms MyString) toLower() MyString {
	return MyString(strings.ToLower(string(ms)))
}

//FindVowels method for MyString
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, v := range ms.toLower() {
		if v == 'a' ||
			v == 'e' ||
			v == 'i' ||
			v == 'o' ||
			v == 'u' {
			vowels = append(vowels, v)
		}
	}
	return vowels
}
