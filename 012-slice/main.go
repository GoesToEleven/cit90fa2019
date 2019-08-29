package main

import "fmt"

func main() {
	// primitive data types
	// first class citizens
	a := "hello world" // string
	b := 4             // int
	c := true          // bool

	// composite literal
	// type{values}

	x := []int{42, 43, 44, 45, 1024}
	fmt.Println(a, b, c)
	fmt.Println(x)

	fmt.Printf("the type of x is %T\n", x)

	for i, v := range x {
		fmt.Println("at index", i, "we have value", v)
	}

	for _, v := range x {
		fmt.Println("we have value", v)
	}

	for i := range x {
		fmt.Println(i)
	}

	y := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	fmt.Println(y[0])
	fmt.Println(y[2:4])
	fmt.Println(y[4:])
	fmt.Println(y[:4])

	y = append(y, 30, 31, 32)
	fmt.Println(y)

	z := []int{35, 36, 37, 38, 39}
	y = append(y, 33, 34)
	y = append(y, z...)
	fmt.Println(y)

	y = y[5:16]
	fmt.Println(y)

	l := []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59}
	m := l[:2]
	n := l[8:]
	l = m
	l = append(l, n...)
	fmt.Println(l)

	g := []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59}
	g = append(g[0:2], g[8:]...)
	fmt.Println(g)

	fmt.Println("len", len(g))
	fmt.Println("cap", cap(g))

	g = append(g, y...)

	fmt.Println(g)

	fmt.Println("len", len(g))
	fmt.Println("cap", cap(g))

	/*
		SLICE
		is a data structure
		made up of
		three values:
		1) len
		2) cap
		3) pointer to an array
	*/

	h := make([]int, 10, 20)
	fmt.Println(len(h))
	fmt.Println(cap(h))
	for i, v := range h {
		fmt.Println(i, v)
	}
	h[0] = 12
	fmt.Println(h)
	// can't do this
	// h[11] = 42
	h = append(h, 42)
	fmt.Println(h)
	fmt.Println(len(h))
	fmt.Println(cap(h))

}
