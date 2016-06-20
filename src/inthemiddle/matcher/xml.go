package matcher

import (
	"gopkg.in/xmlpath.v1"
	"net/http"
)

type XmlRequestMatcher struct {
}

func (m XmlRequestMatcher) Match(req *http.Request, matcher *MatchOption) bool {
    path := xmlpath.MustCompile(matcher.Match.Path)
    root, err := xmlpath.Parse(req.Body)

    if err != nil {
        return false
    }

    if value, ok := path.String(root); ok {
        if value == matcher.Match.Value {
            return true
        }
    }

	return false
}
