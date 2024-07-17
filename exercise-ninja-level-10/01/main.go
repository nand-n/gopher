package main

import "fmt"

/*
using func literal , aka anonymous self-executing func and a buffered chanel

*/

func main() {
	c := make(chan int, 1)
	// go func() {
	// 	c <- 43
	// }()
	fmt.Println(<-c)

}
