package matcher

import (
	"strings"

	httper "github.com/chonla/inthemiddle/httper"
	"launchpad.net/xmlpath"
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
