package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"time"

	"github.com/shazow/communal/internal/httphelper"
	"github.com/shazow/communal/loader"
	"github.com/shazow/communal/loader/hackernews"
	"github.com/shazow/communal/loader/reddit"

	flags "github.com/jessevdk/go-flags"
	"github.com/muesli/termenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

// Version of the binary, assigned during build.
var Version string = "dev"

// Options contains the flag options
type Options struct {
	Pprof   string `long:"pprof" description:"Bind pprof http server for profiling. (Example: localhost:6060)"`
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose logging."`
	Version bool   `long:"version" description:"Print version and exit."`

	Serve struct {
		Bind    string `long:"bind" description:"Address and port to listen on." default:"0.0.0.0:8080"`
		DataDir string `long:"datadir" description:"Path for storing the persistent database."`
	} `command:"serve" description:"Serve a communal frontend."`

	Discover struct {
		Args struct {
			URL string `positional-arg-name:"url" description:"Link to discover"`
		} `positional-args:"yes"`
	} `command:"discover" description:"Crawl metadata about a link."`
}

func exit(code int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(code)
}

func main() {
	options := Options{}
	parser := flags.NewParser(&options, flags.Default)
	p, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		if p == nil {
			fmt.Println(err)
		}
		return
	}

	if options.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	switch len(options.Verbose) {
	case 0:
		logger = logger.Level(zerolog.WarnLevel)
	case 1:
		logger = logger.Level(zerolog.InfoLevel)
	default:
		logger = logger.Level(zerolog.DebugLevel)
	}

	if options.Pprof != "" {
		go func() {
			logger.Debug().Str("bind", options.Pprof).Msg("starting pprof server")
			if err := http.ListenAndServe(options.Pprof, nil); err != nil {
				logger.Error().Err(err).Msg("failed to serve pprof")
			}
		}()
	}

	var cmd string
	if parser.Active != nil {
		cmd = parser.Active.Name
	}
	if err := subcommand(cmd, options); err != nil {
		logger.Error().Err(err).Msgf("failed to run command: %s", cmd)
		return
	}
}

func subcommand(cmd string, options Options) error {
	// Setup signals
	ctx, abort := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func(abort context.CancelFunc) {
		<-sigCh
		logger.Warn().Msg("interrupt received, shutting down")
		abort()
		<-sigCh
		logger.Error().Msg("second interrupt received, panicking")
		panic("aborted")
	}(abort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch cmd {
	case "discover":
		return discover(ctx, options)
	case "serve":
		return errors.New("serve is disabled for now, come back later.")
		//	return serve(ctx, options)
	}

	return fmt.Errorf("unknown command: %s", cmd)
}

var dateLayout = "2006-01-02"

func discover(ctx context.Context, options Options) error {
	link := options.Discover.Args.URL
	logger.Debug().Str("link", link).Msg("discovering")

	client := http.Client{
		Transport: httphelper.TransportWithAgent{
			RoundTripper: http.DefaultTransport,
			UserAgent:    fmt.Sprintf("communal/%s", Version), // TODO: Unhardcode
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

	// TODO: Sort by date?

	p := termenv.ColorProfile()
	formatMeta := func(s string) string {
		return termenv.String(s).Foreground(p.Color("#a9dea1")).String()
	}

	formatLink := func(s string) string {
		return termenv.String(s).Underline().String()
	}

	formatCount := func(i int) string {
		s := "✖" + strconv.Itoa(i)
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

	gProgress, gCtx := errgroup.WithContext(gCtx)
	gProgress.Go(func() error {
		defer close(resChan)
		return g.Wait()
	})

	// Accumulate results
	ordered := []*linkResult{}
	lookup := map[string]*linkResult{}
	count := 0

	for res := range resChan {
		count += 1

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
		fmt.Printf(formatCount(item.Count()))
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
