package exercises

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//setting up the seed for the random number generation
	rand.Seed(time.Now().UnixNano())
}

func genLen() int {
	/* password length
	randomLen := rand.Intn(30-8)+8
	*/
	randomLen := rand.Intn(30-8) + 8
	return randomLen
}

func genPassSlice() []int {
	/* password content
	randomInt := rand.Intn(127-33) + 33
	*/
	result := []int{}
	for i := 0; i < genLen(); i++ {
		result = append(result, rand.Intn(127-33)+33)
	}
	return result
}

func genPass(arr []int) string {
	result := ""
	for _, value := range genPassSlice() {
		// randomy generated int is converted into rune and it is converted
		// into string which represents the ASCII value of generated the int
		result += string(rune(value))
	}
	return result
}

// PasswordGenerator generates a random password everytime.
func PasswordGenerator() {
	fmt.Println(genPass(genPassSlice()))
}
