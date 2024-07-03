package main

import "fmt"

func main() {
	//4!
	fmt.Println(factorial(4))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
