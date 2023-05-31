package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

func main() {
	var conuter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&conuter, 1)
			fmt.Println("Counter\t ", atomic.LoadInt64(&conuter))
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("count:", conuter)
}
