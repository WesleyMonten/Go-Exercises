package main

import (
	"fmt"
	"log"
)

func main() {
	testMyCopyStringSlice()
	var s []string
	s = myAppendStringToSlice(s, "I")
	s = myAppendStringToSlice(s, "love", "programming", "in", "PHP!")
	fmt.Println(s)
}

func myCopyStringSlice(s1 []string, s0 []string) {
	var length int
	if len(s1) < len(s0) {
		length = len(s1)
	} else {
		length = len(s0)
	}

	for i := 0; i < length; i++ {
		s1[i] = s0[i]
	}
}

func myAppendStringToSlice(s1 []string, str ...string) []string {
	lenstr := len(str)
	lenS1 := len(s1)
	var s2 = s1

	if lenS1 < lenstr {
		lenTotal := lenstr + lenS1
		s2 = make([]string, lenTotal)
		copy(s2, s1)
	}

	copy(s2[lenS1:], str)
	return s2
}

func testMyCopyStringSlice() {
	s0 := []string{"To", "be", "or", "not", "to", "be.", "That", "is", "the", "question."}
	var s1 []string
	myCopyStringSlice(s1, s0)
	s2 := make([]string, 4)
	expected := make([]string, 4)
	myCopyStringSlice(s2, s0[0:0])
	copy(expected, s0[0:0])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
	myCopyStringSlice(s2, s0[4:])
	copy(expected, s0[4:])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
}
