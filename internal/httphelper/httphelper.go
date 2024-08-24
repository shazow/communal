package httphelper

import (
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

type FixedRoundTrip string

func (rt FixedRoundTrip) RoundTrip(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(rt))),
		Header:     make(http.Header),
	}, nil
}

// TransportWithAgent is an http.RoundTripper transport wrapper
// that adds default headers and configurations.
type TransportWithAgent struct {
	http.RoundTripper
	UserAgent string
	Logger    zerolog.Logger
}

func (transport TransportWithAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	if transport.UserAgent != "" {
		req.Header.Add("User-Agent", transport.UserAgent)
	}

	transport.Logger.Debug().Str("method", req.Method).Str("url", req.URL.String())

	return transport.RoundTripper.RoundTrip(req)
}
