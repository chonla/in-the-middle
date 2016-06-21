package httper

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	logger "inthemiddle/logger"
)

type Request struct {
	h       RequestHeader
	Payload string
}

type RequestHeader struct {
	method   string
	path     string
	protocol string
	version  string
	headers  []HeaderKeyPair
}

func (r *Request) URL() string {
	protocol := "http://"
	if r.h.protocol == "HTTPS" {
		protocol = "https://"
	}
	host := r.h.Host()
	path := r.h.path
	return protocol + host + path
}

func (h *RequestHeader) Host() string {
	for _, v := range h.headers {
		if strings.ToLower(v.Key) == "host" {
			return v.Value
		}
	}
	return ""
}

func NewRequest(body string) Request {
	r := Request{}
	r.parse(body)
	return r
}

func NewRequestHeader(header string) (h RequestHeader) {
	headers := strings.Split(header, "\r\n")
	if len(headers) < 2 {
		logger.Error(errors.New("Invalid request format."))
		return
	}

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
	r.Payload = payload
}

func (h *RequestHeader) ToString() string {
	buf := fmt.Sprintf("\"%s %s %s/%s\"", h.method, h.path, h.protocol, h.version)
	for _, v := range h.headers {
		buf = buf + " " + v.ToString()
	}
	return buf
}
