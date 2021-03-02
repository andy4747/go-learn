package dsa

import "testing"

/*
Unit Testing of all the methods in stack struct.
*/

func TestAssign(t *testing.T) {
	a := stack{}
	items := []int{1, 2, 3}
	a.Assign(items)

	if CompareSlice(a.items, items) == false {
		t.Errorf("Assignment failed, %v inserted but not assigned", items)
	} else {
		t.Logf("Assignment Succeded, %v inserted in the stack", items)
	}
}

func TestPush(t *testing.T) {
	a := stack{}
	item := 10
	a.Push(item)

	if item != a.items[len(a.items)-1] {
		t.Errorf("Push Failed, %v not pushed in the stack", item)
	} else {
		t.Logf("Push Successful, %v pushed in the stack", item)
	}
}

func TestPop(t *testing.T) {
	a := stack{}
	item := 10
	a.Push(item)
	value := a.Pop()

	if value != item {
		t.Errorf("Popped failed, tried popping the stack but failed %v value returned", value)
	} else {
		t.Logf("Popped Seccuessful, %v returned after popping the stack", value)
	}
}

func TestPeek(t *testing.T) {
	a := stack{}
	a.Push(10)
	a.Push(20)
	value := 20

	if a.Peek() != value {
		t.Errorf("Peek Failed, %v peeked instead of %v", a.Peek(), value)
	} else {
		t.Logf("Peek Successful, %v value peeked", value)
	}
}

// compare items in a slice and returns true both slices have same items
func CompareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if b[index] != value {
			return false
		}
	}
	return true
}
