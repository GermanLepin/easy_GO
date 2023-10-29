package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	incrementor := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()

			value := incrementor
			value++
			incrementor = value

			fmt.Println("number is", value)
			wg.Done()
		}()
	}
	wg.Wait()
}
