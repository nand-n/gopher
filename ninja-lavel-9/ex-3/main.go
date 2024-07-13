package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
using go routines create an increment program
	have a variable to hold the incrementer value
	lounch a bunch of goroutines
		each goroutines should 	\
			read teh incrementer value
				store it in new variable
			yield the processor with runtime.Goshed()
			increment the new variable
			write the value in the new variable back to the incrementer variable
		use waitgroups to wait for all of your goroutines to finish
		the above will create a race conditioon
		prove it is a race conition by using -race flag

*/

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(incrementer)
}
