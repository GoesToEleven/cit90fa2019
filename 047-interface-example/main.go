package main

import "fmt"

type hotdog int
type hamburger string

func (h hotdog) foo() {
	// do somehing
}

func (ha hamburger) foo() {
	// do something
}

type vendor interface {
	foo()
}

func main() {
	var h1 hotdog = 42
	var ha1 hamburger = "yum"

	var w vendor
	w = h1
	w = ha1

	fmt.Println(h1)
	fmt.Println(ha1)
	fmt.Println(w)
}
