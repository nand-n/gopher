package main

import "fmt"

func main() {

	c := make(chan int)
	//send
	go foo(c)
	//recieve
	go bar(c)
	fmt.Println("abut to exit ")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// recieve
func bar(c <-chan int) {
	fmt.Println(<-c)
}
