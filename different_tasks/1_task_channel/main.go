package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()

		for value := range ch {
			fmt.Println("number is", value)
		}
	}()
	wg.Wait()
}
