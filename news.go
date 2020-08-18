package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/fvbock/endless"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

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
