package main

import "fmt"

var z = "happy"

func main() {
	fmt.Println(z)

	z = "So super happy"

	foo()
	bar(z)

	fmt.Println("Program about to end!")
}

func foo() {
	fmt.Println(z)
}

func bar(x string) {
	fmt.Println(x + " happy, happy, happy")
	// I found this file
}
