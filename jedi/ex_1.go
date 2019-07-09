package jedi

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	flavors []string
}

// Ex1 fmt
func Ex1() {
	fmt.Println("Ex1")

	p1 := person{
		first: "mikita",
		last:  "gulag",
		flavors: []string{
			"plombir kak v USSR",
			"шоколадное удовольствие",
		},
	}

	p2 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"Hazelnut"},
	}

	fmt.Println(p1.first, p1.last)

	for _, v := range p1.flavors {
		fmt.Println(v)
	}

	fmt.Println(p2.first, p2.last)

	for _, v := range p2.flavors {
		fmt.Println(v)
	}
}
