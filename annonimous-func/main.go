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
}

func foo() {
	fmt.Println(" foo ran")
}
