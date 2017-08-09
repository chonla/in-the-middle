package matcher

import (
	"regexp"

	httper "github.com/chonla/inthemiddle/httper"
)

type RegexpRequestMatcher struct {
}

func (m RegexpRequestMatcher) Match(req *httper.Request, matcher *MatchOption) bool {
	re, err := regexp.Compile(matcher.Pattern)
	if err != nil {
		return false
	}

	result := re.MatchString(req.URL())

	return result
}
