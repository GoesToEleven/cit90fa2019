package main

import "fmt"

func main() {
	a := 4
	b := 5
	s := addNums(a, b)
	fmt.Printf("%d plus %d is %d\n", a, b, s)
}

func addNums(x, y int) int {
	r := x + y
	return r
}
