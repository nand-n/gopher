package main

import "fmt"

func main() {
	am := map[string]int{
		"Tood":    22,
		"Hennery": 44,
		"Padget":  14,
	}
	fmt.Println("the age of hennery is ", am["Hennery"])

	an := make(map[string]int)
	an["Lukas"] = 28
	delete(an, "Likas")
	an["Steph"] = 45
	fmt.Println(len(an))

	for i, v := range an {
		fmt.Println(i, v)
	}
}
