package badger

import (
	"communal/news/store"

	badger "github.com/dgraph-io/badger/v2"
)

// Open returns a store.Store implementation using Badger as the storage
// driver. The store should be (*badgerStore).Close()'d after use.
func Open(opts badger.Options) (*badgerStore, error) {
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	if err := MigrateLatest(db, opts.Dir); err != nil {
		return nil, err
	}

	s := &badgerStore{
		db: db,
	}

	return s, nil
}

func OpenInMemory() (*badgerStore, error) {
	opts := badger.DefaultOptions("").WithInMemory(true)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	s := &badgerStore{
		db: db,
	}
	return s, nil
}

var _ store.Store = &badgerStore{}

type badgerStore struct {
	db *badger.DB
}

func (s *badgerStore) Close() error {
	return s.db.Close()
}

func (s *badgerStore) AddUser(_ *store.User) error {
	panic("not implemented") // TODO: Implement
}

func (s *badgerStore) GetUserByID(userID store.UserID) (*store.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *badgerStore) AddItem(_ *store.Item) error {
	panic("not implemented") // TODO: Implement
}
