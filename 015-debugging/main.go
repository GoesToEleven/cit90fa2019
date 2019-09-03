package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)

	b := bar(4)
	fmt.Println(b)
}

func foo() int {
	return 42
}

func bar(i int) int {
	return i + foo()
}
