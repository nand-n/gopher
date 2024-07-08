package main

import "fmt"

func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(2, 4.5))
}
