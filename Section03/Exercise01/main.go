package main

import (
	"fmt"

	"github.com/WesleyMonten/Go-Exercises/shared/input"
)

type currency float64

const (
	size = 15
)

func (c currency) String() string {
	s := fmt.Sprintf("€%.2f", float64(c))
	return s
}

func main() {

	var db [size]currency

	for i := 0; i < size; i++ {
		db[i] = currency(input.GetRandFloat(20, 100))
	}

	var total, max, min, avg currency

	for _, v := range db {
		total += v
		if min > v || min == 0.0 {
			min = v
		}
		if max < v {
			max = v
		}
	}
	avg = total / size

	fmt.Printf("Total price: %v\n", total)
	fmt.Printf("Max price: %v\n", max)
	fmt.Printf("Min price: %v\n", min)
	fmt.Printf("Avg price: %v\n", avg)
}
