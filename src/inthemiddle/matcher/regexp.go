package matcher

import (
	"net/http"
    "regexp"
)

type RegexpRequestMatcher struct {
}

func (m RegexpRequestMatcher) Match(req *http.Request, matcher *MatchOption) bool {
    re, err := regexp.Compile(matcher.Pattern)
    if err != nil {
        return false
    }

    result := re.MatchString(req.URL.String())

	return result
}
