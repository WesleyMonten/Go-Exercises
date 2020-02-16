package main

import (
	"math/rand"
	"time"
)

const (
	numTests = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func subSlice(data []int, n, i int) (out []int) {
	l := len(data)
	if l == 0 || 0 > i || 0 > n {
		return
	}
	maxSize := l / n
	if 0 != (l % n) {
		maxSize++
	}
	offset := i * maxSize
	if offset > l {
		return
	}
	actSize := min(maxSize, l-offset)
	out = data[offset : offset+actSize]
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
