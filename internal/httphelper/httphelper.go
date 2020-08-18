package httphelper

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type FixedRoundTrip string

func (rt FixedRoundTrip) RoundTrip(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(string(rt))),
		Header:     make(http.Header),
	}, nil
}

// TransportWithAgent is an http.RoundTripper transport wrapper
// that adds default headers and configurations.
type TransportWithAgent struct {
	http.RoundTripper
	UserAgent string
}

func (transport TransportWithAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	if transport.UserAgent != "" {
		req.Header.Add("User-Agent", transport.UserAgent)
	}
	return transport.RoundTripper.RoundTrip(req)
}
