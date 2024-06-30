package main

import (
	"fmt"
	"sort"
)

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

	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println(medianTwo(n))

}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]/2)
}

func medianTwo(x []float64) float64 {
	//allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)
	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
	}
	return (n[i-1] + n[i]/2)

}
