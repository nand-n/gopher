package main

import "fmt"

func outer() func() int {
	return func() int {
		return 43
	}
}

func main() {
	f := outer()
	fmt.Println(f())
}
