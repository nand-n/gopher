package main

import "fmt"

func addT[T int | float64](a, b T) T {
	return a + b
}

type myNumbers interface {
	int | float64
}

type Number interface {
	int64 | float64
}

func addTInterface[T myNumbers](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(addT(2, 4.5))
	fmt.Println(addTInterface(4.222, 45.330002500022))

	//initalize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	//initalize a map for the float values
	floats := map[string]float64{
		"first":  24.889,
		"second": 45.222,
	}
	fmt.Printf("Non Generic Sums : %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var f float64
	for _, v := range m {
		f += v
	}
	return f
}

// SumIntsOrFloats   sums the values of a map m , it supports both floats and integers
//as map values

func SumIntsOrFloats[k comparable, v int64 | float64](m map[k]v) v {
	var s v
	for _, sum := range m {
		s += sum
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values
func SumNumbers[k comparable, V Number](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
