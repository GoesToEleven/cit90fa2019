package main

import "fmt"
import "github.com/GoesToEleven/cit90fa2019/057-going-through-book/goofy"

var y = 99

const z = 999

type person struct {
	first string
}

func main() {
	fmt.Println(y)

	n := foo()
	fmt.Println(n)

	i := bar()
	fmt.Println(i)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	s := goofy.Moo()
	fmt.Println(s)
}

func foo() int {
	x := 42
	x++
	x *= 2
	return x
}
