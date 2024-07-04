package main

import "fmt"

func main() {
	fmt.Println(add(4, 5))

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(55, 77, subtract)
	fmt.Println(y)
	fmt.Println(paradise("hawai"))
}
func add(a int, b int) int {
	return a + b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func subtract(a int, b int) int {
	return a - b
}

func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
