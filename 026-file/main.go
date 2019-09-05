package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i <= 100; i++ {
		f.WriteString(fmt.Sprintf("I will not talk in class %d\n", i))
	}
}
