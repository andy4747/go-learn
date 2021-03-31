package exercises_test

import (
	"testing"

	"github.com/angeldhakal/go-learn/exercises"
)

func TestPrimeNum(t *testing.T) {
	input := 2
	a := exercises.PrimeNum(input)
	if a != true {
		t.Errorf("Number entered %v is prime but returned %v", input, a)
	} else {
		t.Logf("Number entered %v is prime and successfully returned %v", input, a)
	}
}
