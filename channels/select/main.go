package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	// sender
	go send(100, chan1, chan2, chan3)

	// reciever
	go recive(chan1, chan2, chan3)

	fmt.Println("We are done")
}

func recive(chan1, chan2, chan3 <-chan int) {
	for {
		select {
		case v := <-chan1:
			fmt.Println("element from the chan1", v)
		case v := <-chan2:
			fmt.Println("element from the chan2", v)
		case v, ok := <-chan3:
			if !ok {
				fmt.Println("What?!")
				return
			}

			fmt.Println("element from the chan3", v)
			return
		}
	}
}

func send(number int, chan1, chan2, chan3 chan<- int) {
	for i := 0; i < number; i++ {
		if i%2 == 0 {
			chan1 <- i
		} else if i == number-1 {
			chan3 <- i
		} else {
			chan2 <- i
		}
	}
}
