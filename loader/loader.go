package loader

type Loader interface {
}

type Result interface {
	Title() string
	Submitter() string
	Score() int
	Permalink() string
}

type Results []Result

func Discover(link string) (Results, error) {
	return nil, nil
}

func normalizeLink(link string) string {
	return link
}
