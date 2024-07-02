package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}
func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomthing(h human) {
	h.speak()
}

func main() {
	p1 := secretAgent{
		person: person{
			first: "james",
		},
		ltk: true,
	}
	p2 := person{
		first: "Jenny",
	}
	// p1.speak()
	// p2.speak()
	saySomthing(p1)
	saySomthing(p2)
}
