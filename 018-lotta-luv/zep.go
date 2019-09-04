package main

import "fmt"

var z = "happy"

type hotdog int

var h hotdog = 84

func main() {
	fmt.Println(z)

	z = "So super happy"

	foo()
	bar(z)

	fmt.Println("Program about to end!")

	fmt.Println(h)
	fmt.Printf("%T\n", h)
	fmt.Printf("%d\n", int(h))
	fmt.Printf("%T\n", int(h))
}

func foo() {
	fmt.Println(z)
}

func bar(x string) {
	fmt.Println(x + " happy, happy, happy")
	// I found this file
}
