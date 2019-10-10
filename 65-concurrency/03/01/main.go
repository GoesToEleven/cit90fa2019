package main

import (
	"fmt"
)

func main() {
	const gs = 100
	counter := 0

	c := make(chan int)

	for i := 0; i < gs; i++ {
		go func() {
			c <- 1
		}()
	}

	for i := 0; i < gs; i++ {
		// counter = counter + <-c
		counter += <-c
	}

	fmt.Println("total COUNTER BRO:", counter)
}
