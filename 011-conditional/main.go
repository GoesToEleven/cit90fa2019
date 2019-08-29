package main

import "fmt"

func main() {
	foo()
	bar()
	monkey()
}

func foo() {
	x := 5

	if x == 4 {
		fmt.Println(4)
	} else if x == 5 {
		fmt.Println(5)
	} else if x == 6 {
		fmt.Println(6)
	} else {
		fmt.Println("x was none of those")
	}
}

func bar() {
	x := 4

	switch x {
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println("x was none of those")
	}
}
