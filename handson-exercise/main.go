package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	}
}
