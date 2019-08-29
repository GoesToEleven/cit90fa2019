package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := helpers.Bar()
	fmt.Println("exit")
}
