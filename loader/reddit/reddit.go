package reddit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	iface "github.com/shazow/communal/loader"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"mvdan.cc/xurls/v2"
)

const apiSearch = "https://www.reddit.com/search.json" // ?q=url:$LINK

const userAgent = "cli:github.com/shazow/communal:v1 (by /u/shazow)"

var matchURLs = xurls.Strict()

// Loader for Reddit API. Note that http.Client must provide a ~unique
// UserAgent as that's what reddit segments requests by.
type Loader struct {
	Client http.Client
	Logger zerolog.Logger
}

func (loader *Loader) ID() string {
	return "reddit"
}

// Discover returns more tangential links by crawling submissions and comments.
func (loader *Loader) Discover(ctx context.Context, link string) ([]iface.Result, error) {
	res, err := loader.Search(ctx, link)
	if err != nil {
		return nil, err
	}

	loader.Logger.Debug().Int("hits", len(res)).Msg("discover results")
	return loader.linksFromComments(ctx, res)
}

func (loader *Loader) linksFromComments(ctx context.Context, res []redditListing) ([]iface.Result, error) {
	commentChan := make(chan redditListing)
	g, gCtx := errgroup.WithContext(ctx)

	for _, listing := range res {
		listing := listing // Copy value because we're passing it down the closure
		g.Go(func() error {
			stack, err := listing.Comments(ctx, loader)
			if err != nil {
				return err
			}
			// Traverse comment graph
			// FIXME: Could probably do this with fewer allocations
			for len(stack) != 0 {
				var comment redditListing
				comment, stack = stack[0], stack[1:]
				if replies, err := comment.Replies(); err != nil {
					return err
				} else if len(replies) > 0 {
					stack = append(stack, replies...)
				}
				commentChan <- comment
			}
			return nil
		})
	}

	gProcess, _ := errgroup.WithContext(gCtx)
	gProcess.Go(func() error {
		defer close(commentChan)
		return g.Wait()
	})

	links := []iface.Result{}
	count := 0
	for comment := range commentChan {
		count += 1
		for _, found := range matchURLs.FindAllString(comment.Data.Body, -1) {
			links = append(links, redditLink{
				link:    found,
				comment: comment,
			})
		}
	}

	loader.Logger.Debug().Int("comments", count).Msg("processed comments")
	return links, gProcess.Wait()
}

func (loader *Loader) Search(ctx context.Context, link string) ([]redditListing, error) {
	params := url.Values{}
	params.Set("q", "url:"+link)

	query := apiSearch + "?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		return nil, err
	}

	// TODO: Factor out
	req.Header.Add("User-Agent", userAgent)

	resp, err := loader.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request %q failed with status: %d", req.URL.String(), resp.StatusCode)
	}

	var r redditQueryResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Data.Children, nil
}

type redditQueryResult struct {
	Kind string `json:"kind"` // Listing
	Data struct {
		Children []redditListing `json:"children"`
		After    string          `json:"after"`
		Before   string          `json:"before"`
	} `json:"data"`
}

type redditListing struct {
	Kind string `json:"kind"` // t1 is comment, t3 is story
	Data struct {
		RawReplies json.RawMessage `json:"replies"`

		Archived             bool    `json:"archived"`
		Author               string  `json:"author"`
		AuthorFullname       string  `json:"author_fullname"`
		Body                 string  `json:"body"`
		CreatedUTC           float64 `json:"created_utc"`
		Gilded               int     `json:"gilded"`
		ID                   string  `json:"id"`
		NumComments          int     `json:"num_comments"`
		Over18               bool    `json:"over_18"`
		Permalink            string  `json:"permalink"`
		Score                int     `json:"score"`
		Selftext             string  `json:"selftext"`
		Subreddit            string  `json:"subreddit"`
		SubredditID          string  `json:"subreddit_id"`
		SubredditSubscribers int     `json:"subreddit_subscribers"`
		SubredditType        string  `json:"subreddit_type"`
		Title                string  `json:"title"`
		TotalAwardsReceived  int     `json:"total_awards_received"`
		URL                  string  `json:"url"`
		UpvoteRatio          float64 `json:"upvote_ratio"`
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

func (res redditListing) Replies() ([]redditListing, error) {
	if len(res.Data.RawReplies) == 0 || res.Data.RawReplies[0] != '{' {
		return nil, nil
	}

	var r redditQueryResult
	if err := json.Unmarshal(res.Data.RawReplies, &r); err != nil {
		return nil, fmt.Errorf("failed to parse reddit replies: %s - %v", res.Data.RawReplies, err)
	}
	return r.Data.Children, nil
}

func (res redditListing) Comments(ctx context.Context, loader *Loader) ([]redditListing, error) {
	if res.Kind != "t3" {
		return nil, errors.New("cannot load comments: reddit listing is not a story")
	}

	endpoint := res.Permalink() + ".json"
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := loader.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r []redditQueryResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	// Merge children
	var children []redditListing
	for _, res := range r {
		children = append(children, res.Data.Children...)
	}
	return children, nil
}

type redditLink struct {
	link    string
	comment redditListing
}

func (res redditLink) TimeCreated() time.Time {
	return res.comment.TimeCreated()
}

func (res redditLink) Link() string {
	return res.link
}

func (res redditLink) Submitter() string {
	return res.comment.Submitter()
}

func (res redditLink) Score() int {
	return res.comment.Score()
}

func (res redditLink) Permalink() string {
	return res.comment.Permalink()
}
