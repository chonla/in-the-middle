package cacher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	httper "inthemiddle/httper"
	logger "inthemiddle/logger"
	matcher "inthemiddle/matcher"
)

type Cache []CacheItem

type CacheItem struct {
	Key     string         `json:"name"`
	Matcher RequestMatcher `json:"matcher"`
	//req            *http.Request
	//resp           *http.Response
	currentState   string
	ResponseStates map[string]Responsible `json:"states"`
}

type RequestMatcher struct {
	Pattern string `json:"pattern"`
	Type    string `json:"type,omitempty"`
	Match   string `json:"match,omitempty"`
}

type Responsible struct {
	Return    string `json:"return"`
	NextState string `json:"next_state,omitempty"`
}

func (c *CacheItem) Match(req *http.Request) bool {
	return matcher.Match(req, &matcher.MatchOption{
		Pattern: c.Matcher.Pattern,
		Type: c.Matcher.Type,
		Match: c.Matcher.Match,
	})
}

func (c *CacheItem) dispatchReturn(ret string) string {
	if ret[:7] == "file://" {
		data, err := ioutil.ReadFile(ret[7:])
		if err != nil {
			logger.Error(err)
			return ""
		}
		return string(data)
	}
	return ret
}

func (c *CacheItem) GetResponse() *http.Response {
	r := c.ResponseStates[c.currentState]
	if r.NextState != "" {
		logger.Debug("Switch from state \"" + c.currentState + "\" to state \"" + r.NextState + "\"")
		c.currentState = r.NextState
	}

	ret := c.dispatchReturn(r.Return)

	rbuf := httper.NewResponse(ret)
	code, _ := strconv.Atoi(rbuf.Header.Code)
	vMaj, _ := strconv.Atoi(rbuf.Header.Major)
	vMin, _ := strconv.Atoi(rbuf.Header.Minor)
	h := http.Header{}
	for _, v := range rbuf.Header.Headers {
		h.Add(v.Key, v.Value)
	}
	h.Add("X-Cacher", "In-The-Middle")

	resp := &http.Response{
		Status:     fmt.Sprintf("%s %s", rbuf.Header.Code, rbuf.Header.Message),
		StatusCode: code,
		Proto:      fmt.Sprintf("%s/%s", rbuf.Header.Protocol, rbuf.Header.Version),
		Header:     h,
		ProtoMajor: vMaj,
		ProtoMinor: vMin,
		Body:       ioutil.NopCloser(bytes.NewBufferString(rbuf.Payload)),
	}
	return resp
}
