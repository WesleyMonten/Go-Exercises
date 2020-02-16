package main

import (
	"fmt"

	"github.com/WesleyMonten/Go-Exercises/Section09/Lab1/io"

	"github.com/WesleyMonten/Go-Exercises/Section09/Lab1/ms"
)

type (
	ReadWriter interface {
		io.Reader
		io.Writer
	}
)

func main() {
	var m, _ = ms.NewMemStore(100)
	test(m)
}

func test(rw ReadWriter) {
	var totalBytes int
	n, _ := rw.Write([]byte("Hello, world"))
	totalBytes += n

	n, _ = rw.Write([]byte(". "))
	totalBytes += n

	n, _ = rw.Write([]byte("It is a wonderful day."))
	totalBytes += n

	b := make([]byte, 1)
	_, err := rw.Read(b)
	for err == nil {
		fmt.Print(string(b))
		_, err = rw.Read(b)
	}
	fmt.Println()
}
