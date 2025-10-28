package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	flags "github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
		Loaders []string `long:"loaders" description:"Loaders to use. (Default: all)"`
		Args    struct {
			URL string `positional-arg-name:"url" description:"Link to discover"`
		} `positional-args:"yes"`
	} `command:"discover" description:"Crawl metadata about a link."`
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

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	switch cmd {
	case "discover":
		return discover(ctx, options)
	case "serve":
		return serve(ctx, options)
	}

	return fmt.Errorf("unknown command: %s", cmd)
}
