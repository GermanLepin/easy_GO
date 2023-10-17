package main

// Write a program that uses a channel to signal the cancellation of goroutines. Create multiple
// goroutines that do some work (e.g., count to a large number), and when a cancel signal is sent,
// they should stop their execution gracefully.

// These exercises cover various aspects of goroutines and channels in Go, including basic concurrency,
// channel direction, select statements, fan-out, fan-in patterns, worker pools, pipelines, and
// goroutine cancellation. They can help assess a student's understanding and practical skills in
// concurrent programming with Go.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cancel := make(chan struct{})

	agent := newAgent()
	gs := 100

	agent.wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go agent.worker(i, cancel)
	}

	// sleep one second and send a cansel signal
	time.Sleep(1 * time.Second)
	close(cancel)

	agent.wg.Wait()
	fmt.Println("all workers have finished.")
}

func (a *agent) worker(id int, cancel <-chan struct{}) {
	defer a.wg.Done()
	fmt.Printf("worker %d started\n", id)

	for i := 0; ; i++ {
		select {
		case <-cancel:
			fmt.Printf("Worker %d canceled\n", id)
			return
		default:
			fmt.Printf("Worker %d: Count %d\n", id, i)
		}

		// introduce a small delay to simulate work
		time.Sleep(100 * time.Millisecond)
	}
}

type agent struct {
	wg sync.WaitGroup
}

func newAgent() *agent {
	return &agent{}
}
