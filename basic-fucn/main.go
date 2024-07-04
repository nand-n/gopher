/*
exercise 1 :
Create a func with the identifieer foo that returns anint
Create a func withthe identifier bar that returns an int and a string
Call both funcs
print out their results

exercise 2 :
Create a func with the identifier foo that

	takes in a veriadic parameter of type int
	pass in a value of type []int in to your func (unfurl the []int)
	return the sum of all values of type int passed  ni

Create a func with the identifier bar that

	takes a paramet of type [int]
	returns the sum of all values of type int passed in

# Veriadic function can be called with any number of trailing arguments

Exercise 3 :

	defer  multiple funcion in main
		show that a defferd func runs after the func containging it exits
		determine the ordder in which the multiple defer funcs runs
*/
package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(variadicfunc(xi...))

	defer fmt.Println("0")
}

func foo() int {
	xi := 8
	return xi
}

func bar() (int, string) {
	return 4, "james Bond"
}

func variadicfunc(ii ...int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
