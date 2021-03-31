package dsa

import "errors"

//LinearSearch returns the index if item is found else returns error
func LinearSearch(a []int, item int) (int, error) {
	for i, v := range a {
		if v == item {
			return i, nil
		}
	}
	return 0, errors.New("item not fount")
}
