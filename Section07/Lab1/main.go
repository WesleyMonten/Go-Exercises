package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
	log "github.com/sirupsen/logrus"
)

const (
	replaceChars = "`" + `~!@#$%^&*()-_+=[{]}\|;:'",<.>/?`
)

var (
	res       map[string]int
	waitGroup sync.WaitGroup
	mutex     sync.Mutex
)

func main() {
	if 1 == len(os.Args) {
		log.Fatal("Nothing to process")
	}

	files := os.Args[1:]
	res = make(map[string]int)

	startTime := time.Now()
	log.Info("Processing files...")

	for _, f := range files {
		process(f)
	}

	waitGroup.Wait()
	duration := time.Since(startTime)
	result()
	fmt.Printf("Total time elapsed since start: %v\n", duration)
}

func process(f string) {
	waitGroup.Add(1)

	go func() {
		log.Info("Processing file..")
		reader, err := input.NewFileReader(f)
		if err != nil {
			log.Warnf("Unable to open file: %v", err)
			return
		}

		for reader.Scan() {
			line := reader.Text()

			for _, c := range replaceChars {
				line = strings.Replace(line, string(c), "", -1)
			}

			line = strings.ToLower(line)
			words := strings.Split(line, " ")

			for _, w := range words {
				mutex.Lock()
				res[w] = res[w] + 1
				mutex.Unlock()
			}
		}
		waitGroup.Done()
	}()
}

func result() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range res {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
