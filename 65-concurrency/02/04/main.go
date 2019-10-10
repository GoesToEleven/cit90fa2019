package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("COUNTER FINAL VALUE", counter)
}

/*
ALL OLD SCHOOL
waitgroups
mutex
atomic
*/
