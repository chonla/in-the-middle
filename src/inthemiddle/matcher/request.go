package matcher

import (
	"net/http"
	"net/http/httputil"

	httper "inthemiddle/httper"
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
	Match(*httper.Request, *MatchOption) bool
}

var matchers = map[string]RequestMatcher{}

func Initialize() {
	register("plain", PlainTextRequestMatcher{})
	register("regexp", RegexpRequestMatcher{})
	register("json", JsonRequestMatcher{})
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

	reqBody, err := httputil.DumpRequest(req, true)
	if err != nil {
		return false
	}

	r := httper.NewRequest(string(reqBody))

	return matchers[t].Match(&r, m)
}
