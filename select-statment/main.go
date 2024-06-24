package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Microsecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Microsecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("Value form channel 2", v2)
	}

}
