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
