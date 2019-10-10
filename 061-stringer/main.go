package main

import "fmt"

type person struct {
	first string
}

func (p person) String() string {
	return fmt.Sprintf("hello from %s", p.first)
}

func main() {
	p := person{first: "jenny"}
	fmt.Println(p)
}
