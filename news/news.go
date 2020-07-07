package news

import "communal/news/store"

type News interface {
	AddFeed(url string) error

	Top(n int) ([]*store.Item, error)
}

type news struct {
}
