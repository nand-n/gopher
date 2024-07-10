package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Last string
	Age  int
}

func main() {
	p1 := person{
		Name: "Jems",
		Last: "Bond",
		Age:  32,
	}
	p2 := person{
		Name: "Miss",
		Last: "Moneypenny",
		Age:  27,
	}
	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

}
