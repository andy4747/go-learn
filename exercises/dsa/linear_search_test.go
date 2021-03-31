package dsa

import "testing"

func TestLinearSearch(t *testing.T) {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	item := 70
	result, err := LinearSearch(a, item)
	if err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%v", result)
	}
}
