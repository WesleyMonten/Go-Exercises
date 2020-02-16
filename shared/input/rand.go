package input

import (
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

func GetRandInt(low, max int) int {
	if low >= max {
		return 0
	}
	span := max - low

	n := r.Int()
	return (low + (n % span))
}

func GetRandFloat(low, max int64) float64 {
	if low >= max {
		return 0
	}
	span := max - low

	n := r.Float64() * float64(span)
	return (float64(low) + n)
}
