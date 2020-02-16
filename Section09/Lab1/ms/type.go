package ms

import "errors"

var (
	ErrFull  = errors.New("Storage capacity reached")
	ErrEmpty = errors.New("No more data available")
)

type (
	MemStore struct {
		data     []byte
		capacity uint
		offset   int
	}
)
