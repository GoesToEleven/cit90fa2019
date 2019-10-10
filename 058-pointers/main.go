package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(&x)
	fmt.Printf("%T\n", &x)

	fmt.Println(*&x)
	fmt.Printf("%T\n", *&x)

	fmt.Println(&*&x)
	fmt.Printf("%T\n", &*&x)

	foo(&x)
	fmt.Println(x)
}

func foo(n *int) {
	*n = 999
}
