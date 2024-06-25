package main

import "fmt"

/****
Create a program that has a looop that prints every number form 0 to 99
Modify the program  to run 100 curriculum item

Create a data stucture called a map

*/

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("number : ", i)
	// }

	m := map[string]int{
		"james": 20,
		"Money": 43,
	}

	for k, v := range m {
		fmt.Printf("key %v \t vlaues %v \n", k, v)
	}
}
