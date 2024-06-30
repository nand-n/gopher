package main

import "fmt"

/***
Create a slice of slice of string ([][]stirng) and store the ff data in the multi dimentional slice

"James", "Bond", "Shaken not strined"
"Miss" , "Moneypenny", "I'm 008"
Range over the record and  then range  over the data  in each record

*/

func main() {
	jb := []string{"James", "Bond", "Shaken not strined"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}
	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)
	}
}
