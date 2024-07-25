package main

import "fmt"

/*
CHALLANGE #1:
--Use goroutines and channels to calculate factorial

CHALLANGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?

*/

func main() {
	f := factorial(4)
	for n := range f {
		fmt.Println("Total:", n)
	}

}

// func factorial(n int) int {
// 	total := 1
// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
