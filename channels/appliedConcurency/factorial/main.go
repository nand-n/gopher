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

/*
Tods answer for Why using concurency for factorial
A goroutine , concurency and parrallerism is helpfull when you have a lot of processing to do
if he needed to process thousands of factorial calculation then putting them
into a goroutine and running those calculations concurently and in parrellel
using all the cpu cores on my maching will help those calculation get done
faster
*/
