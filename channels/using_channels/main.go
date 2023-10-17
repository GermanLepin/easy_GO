package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	agent := newAgent()

	const gs = 100000

	agent.wg.Add(gs)
	go sender(c, gs)
	go agent.receiver(c)

	agent.wg.Wait()
	fmt.Println("Alles gute")
}

func sender(c chan<- int, gs int) {
	for i := 0; i < gs; i++ {
		c <- i
	}
	close(c)
}

func (a *agent) receiver(c <-chan int) {
	for i := range c {
		fmt.Println(i)
		a.wg.Done()
	}
}

type agent struct {
	wg sync.WaitGroup
}

func newAgent() *agent {
	return &agent{}
}
