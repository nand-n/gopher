package main

import "fmt"

func main() {
	//array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)
	as := [...]string{"a", "b", "c"}
	fmt.Println(len(as))
	fmt.Println(as)

	//slices
	xs := []string{"Hello", "World"}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Printf("index %v - value  %v \n", i, v)
	}
	fmt.Println(xs[0])
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}

	xs = append(xs, "44", "45")
}
