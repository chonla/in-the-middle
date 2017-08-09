package matcher

import (
	"encoding/json"
	"strings"

	"github.com/jmoiron/jsonq"

	httper "github.com/chonla/inthemiddle/httper"
)

type JsonRequestMatcher struct {
}

func (m JsonRequestMatcher) Match(req *httper.Request, matcher *MatchOption) bool {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(req.Payload), &data)
	jq := jsonq.NewQuery(data)

	path := strings.Split(matcher.Match.Path, "/")
	v, err := jq.String(path...)

	if err != nil {
		return false
	}

	if v == matcher.Match.Value {
		return true
	}

	return false
}
