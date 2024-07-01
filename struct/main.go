package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavour string
}

func main() {
	p1 := person{
		firstName:               "nah",
		lastName:                "deb",
		favoriteIceCreamFlavour: "vanila",
	}
	p2 := person{
		firstName:               "bar",
		lastName:                "d",
		favoriteIceCreamFlavour: "chocolate",
	}
	xs := []person{
		p1, p2,
	}

	for _, v := range xs {
		fmt.Println(v.favoriteIceCreamFlavour)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for _, v := range m {

	}
}
