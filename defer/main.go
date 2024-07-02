package main

import (
	"io"
	"os"
)

func main() {
	defer foo()
	bar()
}

func foo() {

}

func bar() {

}

// returns the file's contents as string
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() //f.clase will run when we're finished

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here
		}

	}
	return string(result), nil // f will closed if wer return here

}
