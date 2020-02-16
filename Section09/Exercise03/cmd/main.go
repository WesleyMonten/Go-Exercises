package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/WesleyMonten/Go-Exercises/Section09/Exercise03/wc"
)

type (
	Writer interface {
		Write(b []byte) (n int, err error)
	}
)

func main() {
	wc1 := new(wc.WriteCounter)
	var w Writer = wc1
	var buf []byte

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Int()%100; i++ {
		w.Write(buf)
	}
	fmt.Println(wc1)
}
