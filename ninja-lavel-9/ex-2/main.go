package main

import "fmt"

/*
This exercise will reinfore our understanding of mehtod sets
	- Create a type person struct
	- attach a method speack to type pointer to person
		- *person
	- create type human interface
		- to implicitly implement the interface , a human must have the speak mehtod
	- create func saySomthing
		- have it take in a human as parameter
		- have it call the speak method
	- show the ff in you r code
		- you CAN pass type *person into saySomething
		- you CANNOT pass the type person into saySomething


*/

type person struct {
	Name string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("helow")
}

func saySomthing(h human) {
	h.speak()
}

func main() {
	p := person{
		Name: "James",
	}

	saySomthing(&p)
}
