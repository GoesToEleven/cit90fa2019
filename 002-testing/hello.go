package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

// foo is an example of a func returning a string
func foo() string {
	return "hello world"
}

// func receiver identifier(parameters) return(s) { code }

// VALUES and TYPE
