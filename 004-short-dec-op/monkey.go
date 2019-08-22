package main

import "fmt"

// var b string = "pkg"
var b = "pkg"

func main() {

	// short declaration operator
	// can't be used outside of code blocks

	// strings are wrapped with "" or ``
	x := "Jack"
	y := 42
	z := 42.47
	a := `here
	is	with 
			backticks`

	// copy a line up or down
	// shift + alt + <arrow key>
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	// multi-cursor editing
	// select next occurence of highlighted
	// command + D
	fmt.Printf("the VALUE stored in x is %v and is of TYPE %T\n", x, x)
	fmt.Printf("the VALUE stored in y is %v and is of TYPE %T\n", y, y)
	fmt.Printf("the VALUE stored in z is %v and is of TYPE %T\n", z, z)
	fmt.Printf("the VALUE stored in a is %v and is of TYPE %T\n", a, a)
	fmt.Printf("the VALUE stored in b is %v and is of TYPE %T\n", b, b)

	fmt.Println("calling foo")
	foo()
}

/*
DECLARE
var x int
declaring a variable
and variables are going to store VALUES of a certain TYPE

ASSIGN
var x int = 42
there is assigning a VALUE of a certain TYPE to a variable

-----
initialize

*/
