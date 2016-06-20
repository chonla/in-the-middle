package matcher

import (
	"strings"

	"gopkg.in/xmlpath.v1"

	httper "inthemiddle/httper"
)

type XmlRequestMatcher struct {
}

func (m XmlRequestMatcher) Match(req *httper.Request, matcher *MatchOption) bool {
    path := xmlpath.MustCompile(matcher.Match.Path)
    root, err := xmlpath.Parse(strings.NewReader(req.Payload))

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
