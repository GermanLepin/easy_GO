package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	incrementor := 0
	const gs = 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()

			v := incrementor
			runtime.Gosched()

			v++
			incrementor = v
			fmt.Println(incrementor)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("end value:", incrementor)
}
