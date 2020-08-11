package reddit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const apiSearch = "https://www.reddit.com/search.json" // ?q=url:$LINK

// Loader for Reddit API. Note that http.Client must provide a ~unique
// UserAgent as that's what reddit segments requests by.
type Loader struct {
	Client http.Client
}

func (loader *Loader) Name() string {
	return "Reddit"
}

func (loader *Loader) Discover(ctx context.Context, link string) ([]redditListing, error) {
	params := url.Values{}
	params.Set("q", "url:"+link)

	req, err := http.NewRequestWithContext(ctx, "GET", apiSearch+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := loader.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r redditQueryResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Data.Children, nil
}

type redditQueryResult struct {
	Data struct {
		Children []redditListing `json:"children"`
	} `json:"data"`
}

type redditListing struct {
	Data struct {
		Subreddit            string  `json:"subreddit"`
		Selftext             string  `json:"selftext"`
		AuthorFullname       string  `json:"author_fullname"`
		Gilded               int     `json:"gilded"`
		Title                string  `json:"title"`
		UpvoteRatio          float64 `json:"upvote_ratio"`
		SubredditType        string  `json:"subreddit_type"`
		TotalAwardsReceived  int     `json:"total_awards_received"`
		Archived             bool    `json:"archived"`
		Over18               bool    `json:"over_18"`
		SubredditID          string  `json:"subreddit_id"`
		ID                   string  `json:"id"`
		Author               string  `json:"author"`
		NumComments          int     `json:"num_comments"`
		Score                int     `json:"score"`
		Permalink            string  `json:"permalink"`
		URL                  string  `json:"url"`
		SubredditSubscribers int     `json:"subreddit_subscribers"`
		CreatedUTC           float64 `json:"created_utc"`
	} `json:"data"`
}

func (res redditListing) TimeCreated() time.Time {
	return time.Unix(int64(res.Data.CreatedUTC), 0)
}

func (res redditListing) Title() string {
	return res.Data.Title
}

func (res redditListing) Submitter() string {
	return res.Data.AuthorFullname
}

func (res redditListing) Score() int {
	return res.Data.Score
}

func (res redditListing) NumComments() int {
	return res.Data.NumComments
}

func (res redditListing) Permalink() string {
	return "https://reddit.com" + res.Data.Permalink
}
