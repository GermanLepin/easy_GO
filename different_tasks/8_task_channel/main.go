package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				chan1 <- i
			}
		}()

	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-chan1)
	}

	fmt.Println("done")

	// ***************** my way *****************//
	chan2 := make(chan int)
	gs := 10
	nums := 10

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < gs; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				for i := 0; i < nums; i++ {
					chan2 <- i
				}
			}()
		}
		wg.Wait()

		close(chan2)
	}()

	for value := range chan2 {
		fmt.Println(value)
	}
}
