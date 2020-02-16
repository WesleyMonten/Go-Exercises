package ms

import "errors"

func (m *MemStore) Read(b []byte) (n int, err error) {
	if m == nil {
		return 0, errors.New("MemStore is empty")
	}

	n = copy(b, m.data[:m.offset])
	read := copy(m.data, m.data[n:m.offset])
	m.offset = read

	if n < len(b) {
		return n, ErrEmpty
	}
	return n, nil
}
