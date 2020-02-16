package ms

import "errors"

func NewMemStore(cap uint) (store *MemStore, err error) {

	store = &MemStore{capacity: cap}
	store.data = make([]byte, cap)

	if store.data == nil {
		return nil, errors.New("Empty data")
	}

	return
}
