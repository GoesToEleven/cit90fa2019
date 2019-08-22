package main

import "fmt"

func main() {
	foo()
	bar()

	x := chip()
	fmt.Println("from chip", x)
	// fmt.Println("from chip", chip())

	dale(43)

	alvin("from alvin", 44)

	y, z := mickey()
	fmt.Println(y, z)

	a, b, c := minney()
	fmt.Println(a, b, c)
}
