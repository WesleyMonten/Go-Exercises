package io

type (
	Writer interface {
		Write(b []byte) (n int, err error)
	}
)
