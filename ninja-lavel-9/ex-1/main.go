package main

import (
	"fmt"
	"sync"
)

/*
In addition to the main goroutine lounch tow additional go routines
	- each additional goroutines should print something out
Use waitgroups to make sure each goroutines finishes your  program exits
*/

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go rout1()
	go route2()

	wg.Wait()

}

func rout1() {
	fmt.Println("Printing some thing out")
	wg.Done()
}
func route2() {
	fmt.Println("Printing something out second time ")
	wg.Done()

}
