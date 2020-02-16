package wc

import (
	"fmt"
)

type (
	WriteCounter uint
)

func (wc *WriteCounter) Write(b []byte) (n int, err error) {

	*wc += WriteCounter(1)

	return len(b), nil
}

func (wc WriteCounter) String() string {
	return fmt.Sprintf("Amount of times the method was called: %v", uint(wc))
}
