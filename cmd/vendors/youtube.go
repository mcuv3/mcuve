package vendors

import "strings"

// Youtube searcher
type Youtube struct{ Domain string }

func NewYoutube() *Youtube {
	return &Youtube{Domain: "https://www.youtube.com"}
}

func (y *Youtube) GetBaseUrl() string {
	return y.Domain
}

func (y *Youtube) KeywordPath(query string) string {
	return "/results?search_query=" + strings.Replace(query, " ", "+", -1)
}
