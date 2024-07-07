package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	x := 42
	fmt.Println(x, "'s memory address is ", &x)
	intDelta(&x)
	fmt.Println(x, "'s updated memory address is ", &x)
	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])

}
