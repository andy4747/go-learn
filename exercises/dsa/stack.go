package dsa

import "fmt"

// Stack data structure in golang
func Stack() {
	a := stack{}
	a.Assign([]int{10, 20, 30, 40, 50})
	a.Push(60)
	a.Push(70)
	a.Push(80)
	a.Push(90)
	a.Push(100)
	fmt.Println(a)
	a.Pop()
	a.Pop()
	a.Pop()
	fmt.Println(a)
	fmt.Println(a.Peek())
}

type stack struct {
	items []int
}

func (s *stack) Assign(arr []int) {
	s.items = append(s.items, arr...)
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() int {
	var a int
	s.items, a = removeIndex(s.items, len(s.items)-1)
	return a
}

func (s stack) Peek() int {
	return s.items[len(s.items)-1]
}

func removeIndex(i []int, index int) ([]int, int) {
	item := i[len(i)-1]

	// to remove item from a slice, create a new slice and append items till
	// the index of item you want to remove, and append all other items after
	// index you want to remove
	i = append(i[:index], i[index+1:]...)
	return i, item
}
