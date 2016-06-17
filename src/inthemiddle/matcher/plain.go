package matcher

import (
	"net/http"
)

type PlainTextRequestMatcher struct {
}

func (m PlainTextRequestMatcher) Match(req *http.Request, matcher *MatchOption) bool {
    result := (req.URL.String() == matcher.Pattern)
	return result
}
