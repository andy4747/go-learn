package exercises

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	return comma(s[:n-1]) + ", " + s[n-1:]
}

func Recursion(){
	a := comma("Angel Dhakal")
	fmt.Println(a)
}
