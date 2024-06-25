/*
**
Question
Create 2 random ints between 0 inclusive and 10 exclusing => [0,10)

	assing them to variables with the identifiers x and y

Print their values
Use an if statment to print the correct descritption

	. x and y are both less than 4
	. x and y are both greater than 6
	. x is greater than or equal to 4 and less than or equal to 6
	. y is not 5
	. none of the previose cased ware met
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)
	fmt.Println(x, y)
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4 ")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is graeter than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previouse cased were met")
	}
}
