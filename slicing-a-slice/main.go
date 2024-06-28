package main

import "fmt"

func main() {
	xs := []int{44, 45, 46, 47, 48, 49, 50}

	//[inlusive:exclusive]
	fmt.Println(xs[1:4])
	//[:exclusive]
	fmt.Println(xs[:2])
	//Deleting from slice
	xs = append(xs[:2], xs[3:]...)
	fmt.Println(xs)

	var matrix [3][3]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}
}
