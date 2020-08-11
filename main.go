package main

import (
	"communal/loader/hackernews"
	"communal/loader/reddit"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/OpenPeeDeeP/xdg"
	flags "github.com/jessevdk/go-flags"
	"github.com/muesli/termenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/fvbock/endless"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
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
		return serve(ctx, options)
	}

	return fmt.Errorf("unknown command: %s", cmd)
}

var dateLayout = "2006-01-02"

func discover(ctx context.Context, options Options) error {
	link := options.Discover.Args.URL
	logger.Debug().Str("link", link).Msg("discovering")

	client := http.Client{
		Transport: &httpTransport{
			RoundTripper: http.DefaultTransport,
			UserAgent:    fmt.Sprintf("communal/%s", Version), // TODO: Unhardcode
		},
	}

	// TODO: Sort by date?
	// TODO: Colorize using termenv

	fmt.Println("Query: ", link)

	p := termenv.ColorProfile()
	formatTitle := func(s string) string {
		return termenv.String(s).Foreground(p.Color("#bf6e01")).String()
	}

	// TODO: Parallelize
	{
		loader := hackernews.Loader{client}
		res, err := loader.Discover(ctx, link)
		if err != nil {
			return err
		}
		logger.Debug().Interface("result", res).Msg("hn")

		out := &strings.Builder{}
		fmt.Fprintf(out, "\n➡️ %s\n", loader.Name())
		for i, item := range res.Hits {
			fmt.Fprintf(out, "  %d. [%s]  %s (by %s with %d points)\n", i+1, item.CreatedAt.Format(dateLayout), formatTitle(item.Title), item.Author, item.Points)
			fmt.Fprintf(out, "     %s\n", item.Permalink())
		}
		fmt.Println(out.String())
	}

	{
		loader := reddit.Loader{client}
		res, err := loader.Discover(ctx, link)
		if err != nil {
			return err
		}
		logger.Debug().Interface("result", res).Msg("reddit")

		out := &strings.Builder{}
		fmt.Fprintf(out, "\n➡️ %s\n", loader.Name())
		for i, item := range res {
			fmt.Fprintf(out, "  %d. [%s]  %s (by %s with %d points)\n", i+1, item.TimeCreated().Format(dateLayout), formatTitle(item.Title()), item.Submitter(), item.Score())
			fmt.Fprintf(out, "     %s\n", item.Permalink())
		}
		fmt.Println(out.String())
	}
	return nil
}

func serve(ctx context.Context, options Options) error {
	// FIXME: This is a placeholder, will be replaced with something real later.
	bind := ":8080"
	if len(os.Args) > 1 {
		bind = os.Args[1]
	}

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Communal.")
	})
	r.Post("/api/share", func(w http.ResponseWriter, r *http.Request) {
		// TODO: CSRF etc
		timestamp := time.Now().UTC()
		link := r.FormValue("link")
		if link == "" {
			fmt.Fprintf(w, "Empty link :(")
			return
		}
		fmt.Printf("-> %s\t%s\t%s\n", timestamp, r.RemoteAddr, link)
		fmt.Fprintf(w, "thanks!")
	})

	fmt.Fprintf(os.Stderr, "listening on %s\n", bind)
	return endless.ListenAndServe(bind, r)
}

// findDataDir returns a valid data dir, will create it if it doesn't
// exist.
func findDataDir(overridePath string) (string, error) {
	path := overridePath
	if path == "" {
		path = xdg.New("communal", "communal").DataHome()
	}
	err := os.MkdirAll(path, 0700)
	return path, err
}

// httpTransport is an http.RoundTripper transport wrapper that adds default
// headers and configurations.
type httpTransport struct {
	http.RoundTripper
	UserAgent string
}

func (transport *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if transport.UserAgent != "" {
		req.Header.Add("User-Agent", transport.UserAgent)
	}
	return transport.RoundTripper.RoundTrip(req)
}
