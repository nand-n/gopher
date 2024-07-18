package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.js")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("const n = [1,2,3,4]; console.log(n.map(a=> a))")
	io.Copy(f, r)
}
