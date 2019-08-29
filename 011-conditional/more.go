package main

import "fmt"

func monkey() {
	x := false

	switch {
	case (3 == 2):
		fmt.Println("not true")
	case (3 == 3):
		fmt.Println("3 == 3 true")
		// fallthrough
	case x:
		fmt.Printf("x is %v \n", x)
		// fallthrough
	default:
		fmt.Println("none of those were true")
	}
}
