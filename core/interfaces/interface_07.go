package interfaces

import "fmt"

/*
It is also possible to compare a type to an interface. If we have a type and if that type implements an interface, it is possible to compare this type with the interface it implements.
*/

//Describer interface
type Describer interface {
	Describe()
}

//Human struct
type Human struct {
	name string
	age  int
}

//Describe func that is declared in Describer interface
func (p Human) Describe() {
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

//Interface7 interface
func Interface7() {
	findType2("Angel Dhakal")
	p := Human{
		name: "Angel Dhakal",
		age:  19,
	}
	findType2(p)
}
