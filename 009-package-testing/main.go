package main

import (
	"fmt"

	"github.com/GoesToEleven/cit90fa2019/009-package-testing/helpers"
)

func main() {
	x := foo()
	fmt.Println(x)

	y := helpers.Bar()
	fmt.Println(y)

	fmt.Println("exit")
}
