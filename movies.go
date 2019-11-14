package main

import (
	"fmt"
)

type actor struct {
	firstName string
	lastName  string
}

type movie struct {
	title       string
	rating      int
	releaseYear int
	actors      []actor
}

func main() {
	var a1, a2, a3, a4, a5, a6, a7, a8, a9 actor
	a1 = actor{firstName: "Linda", lastName: "Hamilton"}
	a2 = actor{"Arnold", "Schwarzenegger"}
	a3 = actor{"Chris", "Evans"}
	a4 = actor{"Jennifer", "Lawrence"}
	a5 = actor{"Margot", "Robbie"}
	a6 = actor{"Jaret", "Leto"}
	a7 = actor{"Ryan", "Reynolds"}
	a8 = actor{"Scarlett", "Johansson"}
	a9 = actor{"Brad", "Pitt"}
	var m1, m2, m3, m4, m5, m6 movie
	m1 = movie{
		title:       "Terminator",
		rating:      6,
		releaseYear: 1990,
		actors:      []actor{a1, a2},
	}
	m2 = movie{
		title:       "Avatar",
		releaseYear: 2012,
		actors:      []actor{a1, a2},
		rating:      4,
	}
	m3 = movie{
		title:       "Jocker",
		releaseYear: 2019,
		actors:      []actor{a1, a2, a3, a4},
		rating:      6,
	}
	m4 = movie{
		title:       "Doctor Sleep",
		releaseYear: 2015,
		actors:      []actor{a1, a3, a6},
		rating:      3,
	}
	m5 = movie{
		title:       "Last Christmas",
		releaseYear: 2019,
		actors:      []actor{a7, a8, a9},
		rating:      4,
	}
	m6 = movie{
		title:       "The Martian",
		releaseYear: 2015,
		actors:      []actor{a2, a5, a6},
		rating:      3,
	}
	var l1, l2, l3 []movie
	var l4 []string
	l1 = []movie{m1, m2, m3}
	l2 = []movie{m4, m5, m6}
	l3 = append(l3, l1...)
	l3 = append(l3, l2...)

	for i := 0; i < len(l3); i++ {
		if l3[i].releaseYear == 2015 {
			l4 = append(l4, l3[i].title)
		} else {
		}
	}
	fmt.Printf("these are movies released in 2015: %\nv", l4)
}
