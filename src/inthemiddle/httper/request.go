package httper

import (
	"fmt"
	"regexp"
	"strings"

	logger "inthemiddle/logger"
)

type Request struct {
	h RequestHeader
	p string
}

type RequestHeader struct {
	method   string
	path     string
	protocol string
	version  string
	headers  []HeaderKeyPair
}

func NewRequest(body string) Request {
	r := Request{}
	r.parse(body)
	return r
}

func NewRequestHeader(header string) (h RequestHeader) {
    headers := strings.Split(header, "\r\n")

	pattern := `([A-Z]+) (.+) (HTTP)/(\d+\.\d+)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		logger.Error(err)
		return
	}
	matches := re.FindStringSubmatch(headers[0])
	h.method = matches[1]
	h.path = matches[2]
	h.protocol = matches[3]
	h.version = matches[4]

    pattern = `([^:]+):\s*(.*)`
    re, err = regexp.Compile(pattern)
    if err != nil {
		logger.Error(err)
		return
	}

    h.headers = []HeaderKeyPair{}
    for _, v := range headers[1:] {
        matches = re.FindStringSubmatch(v)
        h.headers = append(h.headers, HeaderKeyPair{Key: matches[1], Value: matches[2]})
    }

	return
}

func (r *Request) ToString() string {
	return r.h.ToString()
}

func (r *Request) parse(body string) {
	parts := strings.SplitN(body, "\r\n\r\n", 2)
	header := parts[0]
	payload := parts[1]
	r.h = NewRequestHeader(header)
	r.p = payload
}

func (h *RequestHeader) ToString() string {
	buf := fmt.Sprintf("\"%s %s %s/%s\"", h.method, h.path, h.protocol, h.version)
    for _, v := range h.headers {
        buf = buf + " " + v.ToString()
    }
    return buf
}
