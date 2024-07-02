package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("the title of the book is  ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number  ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log From 138", s.String())
}
func main() {
	b := book{
		title: "West withthe night",
	}

	var n count = 43

	fmt.Println(b)
	fmt.Println(n)
	logInfo(b)
	logInfo(n)

}
