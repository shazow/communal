package hackernews

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const apiSearch = "https://hn.algolia.com/api/v1/search"

func normalizeLink(link string) string {
	// Funfact: HN Angolia truncates URLs to 60 chars, so searching for a longer URL yields no results
	if len(link) > 60 {
		link = link[:60]
	}
	// TODO: Strip scheme?
	return link
}

type Loader struct {
	Client http.Client
}

func (loader *Loader) Name() string {
	return "Hacker News"
}

func (loader *Loader) Discover(ctx context.Context, link string) (*hnQueryResult, error) {
	params := url.Values{}
	params.Set("query", normalizeLink(link))
	params.Set("restrictSearchableAttributes", "url")

	req, err := http.NewRequestWithContext(ctx, "GET", apiSearch+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := loader.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r hnQueryResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

type hnHit struct {
	ID        string    `json:"objectID"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Author    string    `json:"author"`
	Points    int       `json:"points"`
	StoryText string    `json:"story_text"`
}

func (hit hnHit) Permalink() string {
	return "https://news.ycombinator.com/item?id=" + hit.ID
}

type hnQueryResult struct {
	Hits []hnHit `json:"hits"`
}
