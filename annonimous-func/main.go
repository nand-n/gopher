package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Annonemous func ran ")
	}()
}

func foo() {
	fmt.Println(" foo ran")
}
