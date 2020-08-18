package hackernews

import (
	iface "communal/loader"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
)

const apiSearch = "https://hn.algolia.com/api/v1/search"

const apiNewsItem = "https://hn.algolia.com/api/v1/items/" // + objectID

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
	Logger zerolog.Logger
}

func (loader *Loader) ID() string {
	return "hackernews"
}

// Discover returns more tangential links by crawling submissions and comments.
func (loader *Loader) Discover(ctx context.Context, link string) ([]iface.Result, error) {
	res, err := loader.Search(ctx, link)
	if err != nil {
		return nil, err
	}

	loader.Logger.Debug().Int("hits", len(res.Hits)).Msg("discover results")
	return loader.linksFromComments(ctx, res)
}

func (loader *Loader) linksFromComments(ctx context.Context, res *hnQueryResult) ([]iface.Result, error) {
	commentChan := make(chan hnComment)
	g, gCtx := errgroup.WithContext(ctx)
	for _, hit := range res.Hits {
		story := hit // Copy value because we're passing it down the closure
		g.Go(func() error {
			stack, err := story.Comments(ctx, loader)
			if err != nil {
				return err
			}
			// Traverse comment graph
			// FIXME: Could probably do this with fewer allocations
			for len(stack) != 0 {
				var comment hnComment
				comment, stack = stack[0], stack[1:]
				stack = append(stack, comment.Children...)
				commentChan <- comment
			}
			return nil
		})
	}

	gProcess, gCtx := errgroup.WithContext(gCtx)
	gProcess.Go(func() error {
		defer close(commentChan)
		return g.Wait()
	})

	var links []iface.Result
	count := 0
	for comment := range commentChan {
		count += 1
		newLinks, err := getLinks(strings.NewReader(comment.Text))
		if err != nil {
			return nil, err
		}
		for _, link := range newLinks {
			links = append(links, hnLink{
				link:    link,
				comment: comment,
			})
		}
	}

	loader.Logger.Debug().Int("comments", count).Msg("processed comments")
	return links, gProcess.Wait()
}

// Search will find submissions of this link.
func (loader *Loader) Search(ctx context.Context, link string) (*hnQueryResult, error) {
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

func (res hnHit) Comments(ctx context.Context, loader *Loader) ([]hnComment, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", apiNewsItem+res.ID, nil)
	if err != nil {
		return nil, err
	}
	resp, err := loader.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r hnNewsItem
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Children, nil
}

type hnLink struct {
	link    string
	comment hnComment
}

func (res hnLink) TimeCreated() time.Time {
	return res.comment.CreatedAt
}

func (res hnLink) Link() string {
	return res.link
}

func (res hnLink) Submitter() string {
	return res.comment.Author
}

func (res hnLink) Score() int {
	return res.comment.Points
}

func (res hnLink) Permalink() string {
	return "https://news.ycombinator.com/item?id=" + strconv.Itoa(res.comment.ID)
}

type hnComment struct {
	ID        int         `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	Author    string      `json:"author"`
	Text      string      `json:"text"`
	Points    int         `json:"points"`
	ParentID  int         `json:"parent_id"`
	StoryID   int         `json:"story_id"`
	Children  []hnComment `json:"children"`
}

type hnNewsItem struct {
	ID        int         `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	Author    string      `json:"author"`
	Title     string      `json:"title"`
	URL       string      `json:"url"`
	Text      string      `json:"text"`
	Points    int         `json:"points"`
	Children  []hnComment `json:"children"`
}

func getLinks(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}
