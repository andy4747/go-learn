package core

import (
	"fmt"
	"math"
)

// Packages demonstrates packages in golang
func Packages() {
	a := 9.1
	b := 9.1

	// prints the ceiling and floor value of a and b
	fmt.Println(math.Ceil(a), math.Floor(b))

	// for working with packages in go, first initialize the go.mod file
	// >>> go mod "package_name"

	// create a folder (same name as the name of the package) a for your packages
	// inside create a go file and make first line the package "folder_name" (which is the name of the package)
}
