package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	amountOfMessages  = 100
	amountOfProducers = 3
	s                 = rand.NewSource(time.Now().Unix())
	r                 = rand.New(s)
	channel           chan string
)

type (
	Producer struct {
		sum   int
		count int
	}
)

func init() {

	flag.IntVar(&amountOfMessages, "msgs", amountOfMessages, "maximum messages per producer, msgs > 0")
	flag.IntVar(&amountOfProducers, "prods", amountOfProducers, "amount of producers, prods > 0")
}

func main() {

	if amountOfMessages < 1 || amountOfProducers < 1 {

		flag.Usage()
		return

	}

	channel = make(chan string, amountOfMessages*amountOfProducers)
	for i := 0; i < amountOfProducers; i++ {

		producer(i + 1)

	}

	close(channel)
	consumer()
}

func consumer() {
	result := make(map[string]Producer)

	for msg := range channel {

		fields := strings.Split(msg, ", num: ")
		value, _ := strconv.ParseInt(fields[1], 10, 64)
		p := result[fields[0]]

		p.sum += int(value)
		p.count++
		result[fields[0]] = p
	}

	sum, count := 0, 0

	for k, v := range result {
		fmt.Printf("%v\n Elements: %v\n total: %v\n", k, v.count, v.sum)
		sum += v.sum
		count += v.count
	}

	fmt.Printf("Overall Total count: %v\n", count)
	fmt.Printf("Overall Total sum: %v\n", sum)
}

func producer(id int) {
	m := (r.Int() % amountOfMessages) + 1

	for i := 0; i < m; i++ {
		channel <- fmt.Sprintf("Producer %v, num: %v", id, r.Int()%20)
	}

}
