package loader

type Loader interface {
}

type Result struct {
}

type Results []Result

func Discover(link string) (Results, error) {
	return nil, nil
}

func normalizeLink(link string) string {
	return link
}
