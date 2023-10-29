package main

import (
	"fmt"
	"sync"
)

func main() {
	agent := newAgent()

	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	agent.wg.Add(2)
	go agent.sender(chan1, chan2)
	go agent.reciver(chan1, chan2, chan3)

	for value := range chan3 {
		fmt.Println("number is", value)
	}

	agent.wg.Wait()
}

func (a *agent) sender(chan1, chan2 chan<- int) {
	defer close(chan1)
	defer close(chan2)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			chan1 <- i
		} else {
			chan2 <- i
		}
	}
	a.wg.Done()
}

func (a *agent) reciver(chan1, chan2 <-chan int, chan3 chan<- int) {
	defer close(chan3)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for value := range chan1 {
			chan3 <- value
		}
		wg.Done()
	}()

	go func() {
		for value := range chan2 {
			chan3 <- value
		}
		wg.Done()
	}()
	wg.Wait()

	a.wg.Done()
}

type agent struct {
	wg sync.WaitGroup
}

func newAgent() *agent {
	return &agent{}
}
