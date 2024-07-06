package main

import "fmt"

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

func main() {
	f := outer()
	fmt.Println(f())

	fmt.Println(printSquare(sauare, 3))
}
