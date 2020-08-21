package loader

import (
	"context"
	"time"
)

type Loader interface {
	ID() string
	Discover(ctx context.Context, link string) ([]Result, error)
}

type Result interface {
	Link() string
	Submitter() string
	Score() int
	Permalink() string
	TimeCreated() time.Time
}

type Results []Result

func Discover(link string) (Results, error) {
	return nil, nil
}
