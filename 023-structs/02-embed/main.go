package main

import "fmt"

type person struct {
	first  string
	last   string
	favnum int
	drinks bool
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first:  "john",
		last:   "doe",
		favnum: 12,
		drinks: false,
	}
	fmt.Println(p1)

	sa1 := secretAgent{
		person: person{
			first:  "james",
			last:   "bond",
			favnum: 7,
			drinks: true,
		},
		ltk: true,
	}
	fmt.Println(sa1)
}
