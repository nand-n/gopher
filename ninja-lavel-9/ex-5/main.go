package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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


		use package atomic

*/

func main() {
	var wg sync.WaitGroup
	var incremente int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&incremente, 1)

			fmt.Println("Incremented by : ", atomic.LoadInt64(&incremente))

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(incremente)
}
