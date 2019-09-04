package main

import "fmt"

var y = "happy"

func main() {
	fmt.Println(y)

	y = "So super happy"

	bar(y)
	foo()

	// fmt.Println("Program about to end!")
}

func foo() {
	fmt.Println(y)
}

func bar(x string) {
	fmt.Println(x + " happy, happy, happy")
}
