package core

import "fmt"

func Strings() {
	a := "angel"
	for _, v := range a {
		fmt.Println(string(byte(v)))
	}
}
