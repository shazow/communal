package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
