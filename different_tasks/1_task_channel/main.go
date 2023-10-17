package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()

		for value := range ch {
			fmt.Println("number is", value)
		}
	}()

	wg.Wait()
}
