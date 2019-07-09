package jedi

import (
	"fmt"
)

// Ex4 fmt
func Ex4() {
	fmt.Println("Ex4")

	person := struct {
		first string
		last  string
		age   rune
	}{
		first: "mikita",
		last:  "gulag",
		age:   23,
	}

	fmt.Println(person)
}
