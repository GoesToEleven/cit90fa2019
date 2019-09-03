package main

import "fmt"

func main() {
	a := foo()
	b := bar()
	c := dog()
	d := mas(4)

	fmt.Println(a, b, c)
	fmt.Println(d)
}

func foo() int {
	return 42
}

func bar() string {
	return "hello"
}

func dog() string {
	return "fido"
}

func mas(i int) int {
	return i + foo()
}
