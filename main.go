package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	flags "github.com/jessevdk/go-flags"
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
}

func exit(code int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(code)
}

func main() {
	options := Options{}
	p, err := flags.NewParser(&options, flags.Default).ParseArgs(os.Args[1:])
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

	if options.Pprof != "" {
		go func() {
			logger.Debug().Str("bind", options.Pprof).Msg("starting pprof server")
			if err := http.ListenAndServe(options.Pprof, nil); err != nil {
				logger.Error().Err(err).Msgf("failed to serve pprof")
			}
		}()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx, options); err != nil {
		logger.Error().Err(err)
	}
}

func run(ctx context.Context, options Options) error {
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
