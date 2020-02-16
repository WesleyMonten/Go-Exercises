package main

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type (
	WriteCounter struct {
		count         int
		amountOfBytes int
	}
)

func (wc *WriteCounter) String() string {
	if wc == nil {
		return ""
	}
	return fmt.Sprintf("\nCount: %v || amount of bytes: %v>\n", wc.count, wc.amountOfBytes)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Invalid amount")
	}

	c := os.Args[1]
	var args []string

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	command := exec.Command(c, args...)
	command.Stdin = os.Stdin

	wc := new(WriteCounter)
	command.Stdout = wc
	err := command.Run()

	if err != nil {
		log.Error(err)
	}

	fmt.Println(wc)
}

func (wc *WriteCounter) Write(b []byte) (n int, err error) {
	if wc == nil {
		return len(b), nil
	}
	wc.count++
	wc.amountOfBytes += len(b)

	return os.Stdout.Write(b)
}
