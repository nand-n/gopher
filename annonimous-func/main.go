package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Annonemous func ran ")
	}()

	func(s string) {
		fmt.Println("this is my name", s)
	}("todd")

	//Returning a funciton

	y := bar()
	y()
}

func foo() {
	fmt.Println(" foo ran")
}

func bar() func() int {
	return func() int {
		return 43
	}
}
