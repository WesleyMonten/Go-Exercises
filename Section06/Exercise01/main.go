package main

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	size      = 100000
	searchers = 6
)

func main() {
	slice, err := createSlice()

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < searchers; i++ {
		data := subSlice(slice, searchers, i)
		go getPrimeNumbers(i+1, data)
	}

	time.Sleep(5 * time.Second)
}

func createSlice() (s []int, err error) {

	s = make([]int, size)

	if s == nil {
		err = errors.New("Error creating slice")
		return
	}

	for i := 0; i < size; i++ {
		s[i] = i + 1
	}
	return
}

func getPrimeNumbers(id int, data []int) {
	i := 0
	for _, val := range data {

		for i = 2; i < val; i++ {

			if (val % i) == 0 {
				break
			}

		}

		if i == val {
			fmt.Println(val)
		}
	}
}
