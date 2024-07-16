package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)
	//recieve
	recieve(eve, odd, quit)

	fmt.Println("about to exit ")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
		q <- 0
	}
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case V := <-e:
			fmt.Println("From the eve channel: ", V)
		case V := <-o:
			fmt.Println("From the odd channel :", V)
		case V := <-q:
			fmt.Println("From the quit channel :", V)
			return
		}
	}
}
