package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	const gs = 100
	var counter int64

	var wg sync.WaitGroup

	c := make(chan int64)

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go func() {
			c <- 1
			wg.Done()
		}()
	}

	go func() {
		for v := range c {
			atomic.AddInt64(&counter, v)
		}
	}()

	wg.Wait()
	close(c)
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("total COUNTER BRO:", atomic.LoadInt64(&counter))
}
