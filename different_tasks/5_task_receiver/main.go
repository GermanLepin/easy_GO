package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hey man!")

}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Jack"}
	saySomething(&p)

	// doesn't work (method speak has pointer receiver)
	//saySomething(p)
}
