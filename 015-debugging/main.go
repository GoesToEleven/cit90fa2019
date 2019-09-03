package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)

	b := sumToFoo(4)
	fmt.Println(b)
}

func foo() int {
	return 42
}

func sumToFoo(i int) int {
	z := i + foo()
	fmt.Println(z)
	return z
}
