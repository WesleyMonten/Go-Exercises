package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
	log "github.com/sirupsen/logrus"
)

const (
	chars = "`" + `~!@#$%^&*()-_+=[{]}\|;:'",<.>/?`
)

var (
	result map[string]int
)

func main() {
	if 1 == len(os.Args) {
		log.Fatal("No files to process")
	}

	inputFiles := os.Args[1:]
	result = make(map[string]int)

	start := time.Now()
	log.Infof("Processing %v files for input", len(inputFiles))
	for _, fn := range inputFiles {
		processFile(fn)
	}

	elapse := time.Since(start)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

func processFile(fn string) {
	log.Infof("Processing file %v", fn)
	reader, err := input.NewFileReader(fn)

	if err != nil {
		log.Warn("Error processing file")
		return
	}

	for reader.Scan() {
		countWords(reader.Text())
	}
}

func countWords(record string) {
	for _, c := range chars {
		record = strings.Replace(record, string(c), "", -1)
	}

	record = strings.ToLower(record)
	words := strings.Split(record, " ")

	for _, w := range words {
		result[w] = result[w] + 1
	}
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
