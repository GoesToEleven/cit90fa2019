package main

import "fmt"

// STACK
// moo
// bar
// foo
// main

// HEAP
// has to be garbage collected

func main() {
	s := foo()
	fmt.Println(*s)

	g()
}

// does this escape to the HEAP?
// no
func g() {
	fmt.Println("IN GGGGG")
	var y int
	fmt.Println(y)

	z := &y
	*z = 42
	fmt.Println(y)
	fmt.Println(*z)
}

func foo() *string {
	s := bar()
	s = fmt.Sprintf("from bar - %s", s)
	// ESCAPES TO THE HEAP
	return &s
}

func bar() string {
	s := moo()
	return fmt.Sprintf("from foo - %s", s)
}

func moo() string {
	return "from moo"
}
