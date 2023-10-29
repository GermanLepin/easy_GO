package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			chan1 <- i
		}
		close(chan1)
	}()

	go func() {
		defer wg.Done()

		for value := range chan1 {
			fmt.Println("number is", value)
		}
	}()
	wg.Wait()

	fmt.Println("DONE")

	// ********************* easy way ***************************** //

	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
