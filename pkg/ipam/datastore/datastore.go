package datastore

import "sync"

type DataStore struct {
	lock sync.Mutex
}

func NewDataStore() (*DataStore, error) {
	return nil, nil
}
