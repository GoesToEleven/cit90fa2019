package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive channel
func receive(c <-chan int) {
	for v := range c {
		fmt.Println("the value received from the channel:", v)
	}
}
// range exits after (A) the chan has been closed and (B) last value has been pulled off

// unbuffered channels can only hold ONE VALUE at a time
