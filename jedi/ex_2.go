package jedi

import (
	"fmt"
)

// Ex2 fmt
func Ex2() {
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Printf("key -> %v \t value -> %v\n", k, v)
	}
}
