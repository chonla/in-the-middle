package matcher

import (
	"net/http"
)

type MatchOption struct {
    Pattern string
	Type    string
	Match   string
}

type RequestMatcher interface {
	Match(*http.Request, *MatchOption) bool
}

var matchers = map[string]RequestMatcher{}

func Initialize() {
    register("plain", PlainTextRequestMatcher{})
}

func register(key string, m RequestMatcher) {
    matchers[key] = m;
}

func Match(req *http.Request, m *MatchOption) bool {
    t := m.Type
    if t == "" {
        t = "plain"
    }
    return matchers[t].Match(req, m)
}
