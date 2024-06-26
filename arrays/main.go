package main

import "fmt"

func main() {
	//array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)
	as := [...]string{"a", "b", "c"}
	fmt.Println(len(as))
	fmt.Println(as)
}
