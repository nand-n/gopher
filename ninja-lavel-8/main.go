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

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
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

	//-----------------Unmarshal

	jsonData := `[{"First":"Nah","Age":25 , "Sayings":["Love me or hate me i dont care , if you are mine you will be mine"]},{"First":"Mercy","Age":23 , "Sayings":["I love you , but you are not giving me the perfect conditions to lvoe you and accept your love , but i have faith in you one day we willbe maried , when you changed your approach."]}]`
	var people []person
	err = json.Unmarshal([]byte(jsonData), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
