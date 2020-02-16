package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
)

var res map[string]int

func main() {
	reader, err := input.NewFileReader("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	res = make(map[string]int)

	var words []string
	var line string

	for err == nil {

		line, err = reader.ReadLine()
		line = strings.TrimSpace(line)
		words = strings.Split(line, " ")
		countWords(words)

	}

	print()
}

func countWords(words []string) {
	for _, word := range words {

		res[word] = res[word] + 1
	}
}

func print() {
	fmt.Printf("%-10s%s\n", "Count", "Word")

	for count, word := range res {
		fmt.Printf("%-10v%s\n", word, count)
	}
}
