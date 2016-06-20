package matcher

import (
	"net/http"

	logger "inthemiddle/logger"
)

type MatchOption struct {
	Pattern string
	Type    string
	Match   MatchItem
}

type MatchItem struct {
	Path  string
	Value string
}

type RequestMatcher interface {
	Match(*http.Request, *MatchOption) bool
}

var matchers = map[string]RequestMatcher{}

func Initialize() {
	register("plain", PlainTextRequestMatcher{})
	register("regexp", RegexpRequestMatcher{})
	register("xml", XmlRequestMatcher{})
}

func register(key string, m RequestMatcher) {
	matchers[key] = m
}

func Match(req *http.Request, m *MatchOption) bool {
	t := m.Type
	if t == "" {
		t = "plain"
	}

	logger.Debug("Match route with " + t)
	return matchers[t].Match(req, m)
}
