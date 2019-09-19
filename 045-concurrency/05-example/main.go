package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	q := 100
	wg.Add(q)
	for i := 0; i < q; i++ {
		go foo(i)
	}
	wg.Wait()
}

func foo(n int) {
	fmt.Println(n)
	wg.Done()
}

// add a counter and increment each time foo prints
// print counter at end of program
// use mutex lock when incrementing counter
// check you have no race condition with -race flag

// read chapter one
