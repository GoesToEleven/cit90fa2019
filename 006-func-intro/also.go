package main

import "fmt"

func foo() {
	fmt.Println("from foo")
}

func bar() {
	fmt.Println("from bar")
}

func chip() int {
	return 42
}

func dale(x int) {
	fmt.Println("from dale", x)
}

func alvin(s string, x int) {
	fmt.Println(s, x)
}

func mickey() (string, int) {
	return "from mickey", 45
}

func minney() (string, int, bool) {
	return "from minney", 46, true
}

/*
a function has this signature:
func receiver identifer(parameters) return(s) {code}

define a func with PARAMETERS
call a func and pass in ARGUMENTS

*/
