package jedi

import (
	"fmt"
)

type vehicle struct {
	doors rune
	color string
}

type truck struct {
	vehicle
	forWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// Ex3 fmt
func Ex3() {
	fmt.Println("Ex3")

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		forWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(t.doors, t.color, t.forWheel)

	fmt.Println(s)
	fmt.Println(s.doors, s.color, s.luxury)
}
