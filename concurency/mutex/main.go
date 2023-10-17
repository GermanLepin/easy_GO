package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()

			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
