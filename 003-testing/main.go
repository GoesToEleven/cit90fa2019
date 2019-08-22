package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

func foo() int {
	return 42
}
