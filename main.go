package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

var index_tmpl = `
<html>
    <body style="width:32em;">

        <p>
            I'm trying to avoid reading news and social media every day; please
            help me by sharing links that you think I will enjoy. I will review
            them at the end of the week.
        </p>

        <form method="post" action="/api/share">

            <label>Link
                <input type="text" name="link" placeholder="https://..." />
            </label>
            <input type="submit" value="Check this out" />

        </form>

    </body>
</html>
`

type APIResult struct {
	Code     int                    `json:"code"`
	Status   string                 `json:"status"`
	Messages []string               `json:"messages"`
	Result   map[string]interface{} `json:"result"`
}

func main() {
	bind := ":8080"
	if len(os.Args) > 1 {
		bind = os.Args[1]
	}

	r := chi.NewRouter()

	//r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, index_tmpl)
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
	endless.ListenAndServe(bind, r)
}
