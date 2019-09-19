package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	fmt.Println("START CPUs:", runtime.NumCPU())
	fmt.Println("START Goroutines:", runtime.NumGoroutine())

	launchRoutines()

	wg.Wait()
	fmt.Println("FINISH Goroutines:", runtime.NumGoroutine())
	fmt.Println("FINISH count:", counter)
}

func launchRoutines() {
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Second)
			runtime.Gosched()
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("WORKING Goroutines:", runtime.NumGoroutine())
	}
}
