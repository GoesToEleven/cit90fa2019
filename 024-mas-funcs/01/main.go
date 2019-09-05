package main

import "fmt"

func main() {
	xi := []int{1, 2, 4, 8, 16}
	s := foo(xi)
	fmt.Println(s)

	t := bar(42, 43, 44, 99, 1024)
	fmt.Println(t)

	u := bar(xi...)
	fmt.Println(u)
}

// we can pass a slice of int into a func
func foo(ii []int) int {
	var sum int
	for _, v := range ii {
		fmt.Println("adding to total:", v)
		sum += v
	}
	return sum
}

func bar(ii ...int) int {
	var sum int

	fmt.Println("bar running", ii)
	fmt.Printf("%T\n", ii)

	for _, v := range ii {
		fmt.Println("adding to total:", v)
		sum += v
	}
	return sum
}
