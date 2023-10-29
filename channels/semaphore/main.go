package main

import (
	"fmt"
	"time"
)

// "The problem of philosophers having lunch". It is formulated as follows:
// five philosophers sit at a round table; each of them has a plate with
// spaghetti in front of him; to the left of the plate, each of them has a
// fork. Thus, there are five plates and five forks on the table. Each
// philosopher can either eat or ponder. To eat, two forks are needed. The
// forks cannot be taken away. The task is to ensure that each philosopher
// can both think and eat (i.e., it is necessary to correctly distribute
// control over resources).

func main() {
	semaphore := make(chan int, 3)
	done := make(chan struct{})
	i := 0

	go func() {
		for ; ; i++ {
			semaphore <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		time.Sleep(time.Millisecond * 1000)
		done <- struct{}{}
	}()

	msg := 0
L:
	for {
		select {
		case msg = <-semaphore:
			fmt.Println(msg)
		case <-done:
			fmt.Println("done")
			break L
		default:
			if msg >= 20 {
				break L
			}
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 200)
		}
	}

	fmt.Println("success")
}
