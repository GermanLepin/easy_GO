package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		q <- 1

		close(ch)
	}()

	return ch
}

func receive(c, q <-chan int) {
	for {
		select {
		case value := <-c:
			fmt.Println("ooooo this is a value from c channel", value)
		case value := <-q:
			fmt.Println("ooooo this is a value from q channel", value)
			return
		}
	}
}
