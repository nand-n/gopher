/*
Create a func with the identifieer foo that returns anint
Create a func withthe identifier bar that returns an int and a string
Call both funcs
print out their results
*/
package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	xi := 8
	return xi
}

func bar() (int, string) {
	return 4, "james Bond"
}
