package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func main() {
	x := 42
	fmt.Println(x, "'s memory address is ", &x)
	intDelta(&x)
	fmt.Println(x, "'s updated memory address is ", &x)

}
