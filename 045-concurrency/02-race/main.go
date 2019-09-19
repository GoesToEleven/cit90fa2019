package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("START CPUs:", runtime.NumCPU())
	fmt.Println("START Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Second)
			runtime.Gosched()
			counter++
			wg.Done()
		}()
		fmt.Println("WORKING Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("FINISH Goroutines:", runtime.NumGoroutine())
	fmt.Println("FINISH count:", counter)
}
