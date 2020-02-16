package main

import (
	"fmt"
)

func main() {
	var sumOddNumbers int
	for x := 1; x <= 10001; x++ {
		if x%2 == 1 {
			sumOddNumbers += x
		}
	}

	fmt.Printf("Sum of ODD numbers from 1 to 10001: %v\n", sumOddNumbers)

	var counter, sumNumbers int
	for x := 1; x <= 10001; x++ {
		switch x + 1 {
		case 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910:
			continue
		}
		counter++
		sumNumbers += x
	}
	fmt.Printf("Average: %v\n", sumNumbers/counter)
}
