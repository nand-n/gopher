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

	//Closure funciton
	f := incrementor()
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3

}
func incrementor() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func foo() {
	fmt.Println(" foo ran")
}

func bar() func() int {
	return func() int {
		return 43
	}
}
