package main

import (
	"io"
	"os"
)

var (
	digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func main() {
	io.WriteString(os.Stdout, intToStr(2010)+"\n")
	io.WriteString(os.Stdout, intToStr(1971/07/07)+"\n")
	io.WriteString(os.Stdout, intToStr(29093423)+"\n")
	io.WriteString(os.Stdout, f64ToStr(9.15)+"\n")
	io.WriteString(os.Stdout, f64ToStr(11.04)+"\n")
}

func intToStr(i int) (s string) {

	for {

		r := i % 10
		i /= 10
		s = digits[r] + s

		if i == 0 {

			return
		}
	}
}

func f64ToStr(f float64) (s string) {

	w := int(f)
	s = intToStr(w)
	s += "."
	d := f - float64(w)

	for {
		t := int(d * 10)
		s += intToStr(t)
		d = (d * 10) - float64(t)

		if d < 0.001 {

			return
		}
	}
}
