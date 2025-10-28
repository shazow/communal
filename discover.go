package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/shazow/communal/internal/httphelper"
	"github.com/shazow/communal/loader"
	"github.com/shazow/communal/loader/hackernews"
	"github.com/shazow/communal/loader/reddit"

	"github.com/muesli/termenv"
	"golang.org/x/sync/errgroup"
)

var dateLayout = "2006-01-02"

func discover(ctx context.Context, options Options) error {
	link := options.Discover.Args.URL
	logger.Debug().Str("link", link).Msg("discovering")

	client := http.Client{
		Transport: httphelper.TransportWithAgent{
			RoundTripper: http.DefaultTransport,
			UserAgent:    fmt.Sprintf("cli:github.com/shazow/communal@%s (by /u/shazow)", Version), // TODO: Unhardcode
			Logger:       logger,
		},
	}

	loaders := map[string]loader.Loader{
		"hackernews": &hackernews.Loader{
			Client: client,
			Logger: logger.With().Str("loader", "hackernews").Logger(),
		},
		"reddit": &reddit.Loader{
			Client: client,
			Logger: logger.With().Str("loader", "reddit").Logger(),
		},
	}

	if options.Discover.Loaders != nil {
		filteredLoaders := map[string]loader.Loader{}
		for _, name := range options.Discover.Loaders {
			if loader, ok := loaders[name]; ok {
				filteredLoaders[name] = loader
			} else {
				logger.Warn().Str("loader", name).Msg("unknown loader requested, skipping")
			}
		}
		logger.Debug().Int("available", len(loaders)).Int("filtered", len(filteredLoaders)).Msg("filtered loaders")
		loaders = filteredLoaders
	}

	// TODO: Sort by date?

	p := termenv.ColorProfile()
	formatMeta := func(s string) string {
		return termenv.String(s).Foreground(p.Color("#a9dea1")).String()
	}

	formatLink := func(s string) string {
		return termenv.String(s).Underline().String()
	}

	formatCount := func(i int) string {
		s := "✖ " + strconv.Itoa(i)
		switch i {
		case 1:
			return s
		case 2:
			return termenv.String(s).Foreground(p.Color("#bc9923")).String()
		case 3:
			return termenv.String(s).Foreground(p.Color("#bc7123")).String()
		default:
			return termenv.String(s).Foreground(p.Color("#bc4523")).String()
		}
	}

	resChan := make(chan loader.Result)
	g, gCtx := errgroup.WithContext(ctx)

	for _, loader := range loaders {
		loader := loader // Copy for closure
		g.Go(func() error {
			r, err := loader.Discover(ctx, link)
			if err != nil {
				return err
			}
			for _, res := range r {
				resChan <- res
			}
			return nil
		})
	}

	gProgress, _ := errgroup.WithContext(gCtx)
	gProgress.Go(func() error {
		defer close(resChan)
		return g.Wait()
	})

	// Accumulate results
	ordered := []*linkResult{}
	lookup := map[string]*linkResult{}
	count := 0

	for res := range resChan {
		count++

		if entry, ok := lookup[res.Link()]; ok {
			entry.Add(res)
		} else {
			entry := &linkResult{
				link: res.Link(),
			}
			entry.Add(res)
			lookup[res.Link()] = entry
			ordered = append(ordered, entry)
		}
	}

	if err := gProgress.Wait(); err != nil {
		return err
	}

	sort.Slice(ordered, func(i, j int) bool {
		if a, b := ordered[i].Score(), ordered[j].Score(); a != b {
			return a > b
		} else if a, b := ordered[i].TimeCreated(), ordered[j].TimeCreated(); a != b {
			return a.Before(b)
		}
		return false
	})

	logger.Debug().Int("total", count).Int("deduped", len(ordered)).Msg("result summary")

	for _, item := range ordered {
		fmt.Printf("%s ", formatLink(item.Link()))
		fmt.Print(formatCount(item.Count()))
		fmt.Printf(formatMeta(" on %s")+"\n", item.TimeCreated().Format(dateLayout))
	}

	return nil
}

type linkResult struct {
	link        string
	timeCreated time.Time
	results     []loader.Result
}

func (res *linkResult) Count() int {
	return len(res.results)
}

func (res *linkResult) Add(r loader.Result) {
	res.results = append(res.results, r)

	if res.timeCreated.IsZero() {
		res.timeCreated = r.TimeCreated()
	} else if res.timeCreated.After(r.TimeCreated()) {
		res.timeCreated = r.TimeCreated()
	}
}

func (res *linkResult) Submitter() string {
	return fmt.Sprintf("%d people", len(res.results))
}

func (res *linkResult) Score() int {
	return len(res.results)
}

func (res *linkResult) Link() string {
	return res.link
}

func (res *linkResult) TimeCreated() time.Time {
	return res.timeCreated
}
