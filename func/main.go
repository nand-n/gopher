package main

import "fmt"

func main() {
	foo()
	s, n := dogYears("Jack", 5)
	fmt.Println(s, n)
}

func foo() {
	fmt.Println("I am foo")
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is  this old in dog years."), age
}
