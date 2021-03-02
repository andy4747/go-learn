package core

import (
	"fmt"
)

// Student struct
type Student struct {
	Name   string
	Grades []int
	Age    int
}

func (s Student) getAge() int {
	return s.Age
}

// AverageGrades returns the average grade of a student
func (s Student) AverageGrades() float64 {
	total := 0
	for _, value := range s.Grades {
		total += value
	}
	return float64(total) / float64(len(s.Grades))
}

func (s *Student) setAge(age int) {
	s.Age = age
}

// Methods is a function that contains methods of a struct
func Methods() {
	fmt.Println("Methods")
	s1 := Student{"Angel Dhakal", []int{89, 87, 91, 94}, 19}
	fmt.Println(s1.AverageGrades())
}
