package ms

import "errors"

func (m *MemStore) Write(b []byte) (n int, err error) {
	if m == nil {
		return 0, errors.New("Write called on nil value")
	}

	n = copy(m.data[m.offset:], b)
	m.offset += n

	if n < len(b) {
		return n, ErrFull
	}
	return n, nil
}
