package main

import "fmt"

func addT[T int | float64](a, b T) T {
	return a + b
}

type myNumbers interface {
	int | float64
}

func addTInterface[T myNumbers](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(addT(2, 4.5))
	fmt.Println(addTInterface(4.222, 45.330002500022))
}
