package exercises

import "math"

// PrimeNum takes num of type int as an input
// and returns true if num is prime else false
func PrimeNum(num int) bool {
	if num > 1 {
		a := math.Sqrt(float64(num))
		for i := 2; i <= int(a); i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
