package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is ", d.first, "and i'm walking")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is ", d.first, "and i'm runing `")
}

type youngin interface {
	walk()
	run()
}

func younginRun(y youngin) {
	y.run()
}

type person struct {
	first string
}

func main() {
	d1 := dog{"henery"}
	d1.walk()
	d1.run()
	younginRun(&d1)
	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	p1 := person{
		first: "Jenny",
	}

	p1 = changeName(p1, "Jen")
	changeNamePtr(&p, "JMP")
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}
