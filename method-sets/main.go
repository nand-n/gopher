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

func main() {
	d1 := dog{"henery"}
	d1.walk()
	d1.run()
	younginRun(&d1)
	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
}
