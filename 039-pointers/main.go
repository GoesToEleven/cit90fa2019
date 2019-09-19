package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	// '&' shows the memory location where a VALUE is stored
	fmt.Println("x is stored at this location", &x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	// a *int is a TYPE
	var a *int
	a = &x
	fmt.Println(a)

	// you can DEREFERENCE a pointer with the '*' operator
	fmt.Println("here's the value that's stored in the memory location", *a)
}
