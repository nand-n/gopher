package main

func main() {
	cs := make(chan int)
	go func() {
		cs <- 42
	}()

}
