package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// UserInput is a function that takes two users input from user and adds them
func UserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	// taking first user input
	fmt.Print("Enter first number: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // C onverting string input to int

	// taking second user input
	fmt.Print("Enter second number: ")
	scanner.Scan()
	input2, _ := strconv.ParseInt(scanner.Text(), 10, 64) // C onverting string input to int

	// adding two user inputs
	result := input + input2
	fmt.Println(result)
}
