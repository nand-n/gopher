package main

import "fmt"

func main() {
	fmt.Println(Add(4, 5))

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(55, 77, subtract)
	fmt.Println(y)
}
func add(a int, b int) int {
	return a + b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func subtract(a int, b int) int {
	return a - b
}
