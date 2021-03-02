package core

import "fmt"

// Point struct
type Point struct {
	x int
	y int
}

func changeX(pt *Point) {
	pt.x = 20
}

// Struct defines struct in golang
func Struct() {
	p1 := Point{1, 2}
	fmt.Println(p1)

	changeX(&p1)
	fmt.Println(p1)

}
