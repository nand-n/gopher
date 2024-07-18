package main

import "fmt"

/*
Write a program that launches 10 go routines
	each goroutines adds 10 numbers to a channel
put the numbers off the channle and print them.
*/

func main() {
	c := make(chan int)
	for i := 0; i <= 10; i++ {
		go func() {
			for i := 0; i <= 10; i++ {
				c <- i
			}
			// close(c)
		}()

	}
	for i := 0; i < 100; i++ {

		fmt.Println(<-c)
	}
}
