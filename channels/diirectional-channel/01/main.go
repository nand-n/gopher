package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Printf("%T\n", c)

	cr := make(<-chan int) //reciever
	cs := make(chan<- int) // send

	fmt.Println(<-cr)
	fmt.Println(cr)
	fmt.Println("-------")
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

}
