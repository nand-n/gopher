package main

import (
	"fmt"
	"math"
)

func outer() func() int {
	return func() int {
		return 43
	}
}

// creating a calll back function
func sauare(n int) int {
	return n * n
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

// blocked scope function \
func powenator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

func main() {
	f := outer()
	fmt.Println(f())
	fmt.Println(printSquare(sauare, 3))
	x := powenator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
