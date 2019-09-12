package main

import "fmt"

type hotdog int

var nyStand hotdog

func (h hotdog) String() string {
	return fmt.Sprintf("There are %d stands in New York\n", h)
}

func (h hotdog) Another() {
	fmt.Println("here's another example", h)
}

func main() {
	nyStand = 4
	fmt.Println(nyStand)

	z := nyStand.String()
	fmt.Println("the assigned z:", z)

	nyStand.Another()
}

func foo(x int) {
	fmt.Println(x)
}
