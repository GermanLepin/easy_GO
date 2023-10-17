package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()

			v++
			counter = v
		}()
		fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	}

	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
