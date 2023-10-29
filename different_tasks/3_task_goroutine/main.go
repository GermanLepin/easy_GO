package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	agent := newAgent()

	chan1 := make(chan int)
	chan2 := make(chan int)

	agent.wg.Add(3)
	go agent.generateRandomNumbers(chan1)
	go agent.calculateTheSquare(chan1, chan2)
	go agent.printTheSquare(chan2)

	agent.wg.Wait()
}

func (a *agent) generateRandomNumbers(ch chan<- int) {
	defer a.wg.Done()

	for i := 0; i < 10; i++ {
		num := rand.Intn(1000)
		ch <- num
	}
	close(ch)
}

func (a *agent) calculateTheSquare(ch1 <-chan int, ch2 chan<- int) {
	defer a.wg.Done()

	for value := range ch1 {
		i := value * value

		ch2 <- i
	}
	close(ch2)
}

func (a *agent) printTheSquare(ch <-chan int) {
	defer a.wg.Done()

	for value := range ch {
		fmt.Println("number is", value)
	}
}

type agent struct {
	wg sync.WaitGroup
}

func newAgent() *agent {
	return &agent{}
}
