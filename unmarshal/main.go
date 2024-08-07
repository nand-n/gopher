package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var jsonBlob = []byte(`
	[ { "Name":"Platypus" , "Order":"Monotremata"},
	  {"Name":"Quoll" , "Orders":"Dasynormusta" }
	]
	`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("%+v", animals)
}
