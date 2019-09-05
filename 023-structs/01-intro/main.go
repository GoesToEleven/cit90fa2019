package main

import "fmt"

type person struct {
	first  string
	last   string
	dob    string
	favnum int
	drinks bool
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		dob:    "01/01/1987",
		favnum: 7,
		drinks: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.dob)
	fmt.Println(p1.favnum)
	fmt.Println(p1.drinks)

	// reminder about slice
	xi := []int{42, 43, 44, 99, 23423}
	
	// reminder about map
	m := map[int]string{
		7: "james",
		8: "jenny",
		9: "jack",
	}

	fmt.Println(xi)
	fmt.Println(m)
}
