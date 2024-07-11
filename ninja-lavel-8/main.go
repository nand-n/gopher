package main

import (
	"encoding/json"
	"fmt"
)

/*
Marshal and unmarshal it
*/
type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Nah",
		Age:   25,
	}
	u2 := user{
		First: "Mercy",
		Age:   23,
	}
	users := []user{u1, u2}

	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
