package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mu sync.Mutex

func main() {
	fmt.Println(runtime.NumCPU())

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Second)
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine())
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("COUNTER FINAL VALUE", counter)
}
