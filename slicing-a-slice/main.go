package main

import "fmt"

func main() {
	xs := []int{44, 45, 46, 47, 48, 49, 50}

	//[inlusive:exclusive]
	fmt.Println(xs[1:4])
	//[:exclusive]
	fmt.Println(xs[:2])

}
