package core

import (
	"fmt"
)

// DataTypes is explanation of golang data types
func DataTypes() {
	// integer => default value will be 0 even if it is uint or int

	// uint8 => unsigned 8 bit integer
	// uint16 => unsigned 16 bit integer
	// uint32 => unsigned 32 bit integer
	// uint64 => unsigned 64 bit integeer

	// int8 => signed 8 bit integer
	// int16 => signed 16 bit integer
	// int32 => signed 32 bit integer
	// int64 => signed 64 bit integer

	// float => default value will be 0

	// float32 => 32 bit floating point number
	// float64 => 64 bit floating point number

	// byte => byte data type represents ASCII characters | default value is 0
	var myLetter = 'R' // Type inferred as rune which is the default type for character values
	fmt.Println(myLetter)

	// rune => rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format | default value is 0
	var anotherLetter byte = 'B'
	fmt.Println(anotherLetter)

	// string => default value will be ""

	// boolean (true or false) => default value is false

	// complex => default value is

	// array => default value of an array is an array of default values of its type.

	// struct => Lets suppose there is a struct "sample" with two field. One is of "int" type and the other is of the "bool" type. We create an instance of this struct and when we print it, the output is "{0 false}"

	// map => default value is nil

	// channel => default value is nil

	// interface => default value is nil

	// slice => default value is nil

	// pointer => default value is nil

	// function => default value is nil

}
