package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	f, err := os.Create("Output.txt")

	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer f.Close()

	s := []byte("Hello pohers!")

	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("error %s", err)
	}

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers ")
	fmt.Println(b.String())
}
