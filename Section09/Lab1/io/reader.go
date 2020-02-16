package io

type (
	Reader interface {
		Read(b []byte) (n int, err error)
	}
)
