package main

import "fmt"

type person struct {
	word string
}

func (p *person) speak() string {
	return p.word
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println("what? ", h.speak())
}

func main() {
	x := person{"Hai Apakabar"}
	saySomething(&x)
	fmt.Println("Dua dong: ", x.speak())
}
