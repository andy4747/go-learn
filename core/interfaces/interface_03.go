package interfaces

import "fmt"

/*
An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.
*/

//Worker interface
type Worker interface {
	Work()
}

//Person struct
type Person struct {
	name string
}

//Work struct
func (p Person) Work() {
	fmt.Println(p.name, " is working.")
}

func describe(w Worker) {
	fmt.Printf("Interface type { %T } and value { %v }\n", w, w)
}

//Interface3 function
func Interface3() {
	p := Person{
		name: "Angel Dhakal",
	}
	var w Worker = p
	describe(w)
	w.Work()
}
