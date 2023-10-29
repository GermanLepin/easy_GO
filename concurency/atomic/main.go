package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	var (
		counter int64 // atomic needs
		wg      sync.WaitGroup
	)

	const gs = 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter:\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
