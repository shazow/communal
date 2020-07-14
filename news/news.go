package news

import (
	"communal/news/store"
	"errors"
)

type News interface {
	AddFeed(url string) error

	Top(n int) ([]*store.Item, error)
}

func New(id string) (*news, error) {
	return nil, errors.New("not implemented")
}

type news struct {
	store.Store

	ID store.NewsID
}
