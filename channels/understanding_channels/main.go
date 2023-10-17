package main

import "fmt"

func main() {
	bchan()
	unbchan()
}

func bchan() {
	c := make(chan int, 1)

	c <- 44

	fmt.Println(<-c)

}

func unbchan() {
	c := make(chan int)

	go func() {
		c <- 44
	}()

	fmt.Println(<-c)
}
