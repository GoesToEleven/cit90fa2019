package main

import "fmt"

func main() {
	foo()
	bar()
	monkey()
}

func foo() {
	i := 100
	for {
		fmt.Println(i)
		i--
		if i == 0 {
			break
		}
	}
}

func bar() {
	// for init; cond; post {}

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func monkey() {
	i := 0
	for i <= 100 {
		fmt.Println(i)
		i++
	}
}
