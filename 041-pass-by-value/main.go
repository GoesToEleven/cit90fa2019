package main

import "fmt"

func main() {
	n := 42
	fmt.Println("before foo:", n)
	foo(n)
	fmt.Println("after foo:", n)
}

func foo(x int) {
	x++
	fmt.Println("in foo:", x)
}
