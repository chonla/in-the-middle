package httper

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	logger "github.com/chonla/inthemiddle/logger"
)

type Response struct {
	Header  ResponseHeader
	Payload string
}

type ResponseHeader struct {
	Protocol string
	Version  string
	Major    string
	Minor    string
	Code     string
	Message  string
	Headers  []HeaderKeyPair
}

func NewResponse(body string) Response {
	r := Response{}
	r.parse(body)
	return r
}

func (r *Response) ToString() string {
	return r.Header.ToString()
}

func (r *Response) parse(body string) {
	parts := strings.SplitN(body, "\r\n\r\n", 2)
	if len(parts) < 2 {
		logger.Error(errors.New("Invalid response format."))
		return
	}

	header := parts[0]
	payload := parts[1]
	r.Header = NewResponseHeader(header)
	r.Payload = payload
}

func NewResponseHeader(header string) (h ResponseHeader) {
	headers := strings.Split(header, "\r\n")

	pattern := `(HTTP)/(\d+)\.(\d+) (\d+) (.+)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		logger.Error(err)
		return
	}
	matches := re.FindStringSubmatch(headers[0])

	h.Protocol = matches[1]
	h.Version = matches[2] + "." + matches[3]
	h.Major = matches[2]
	h.Minor = matches[3]
	h.Code = matches[4]
	h.Message = matches[5]

	pattern = `([^:]+):\s*(.*)`
	re, err = regexp.Compile(pattern)
	if err != nil {
		logger.Error(err)
		return
	}

	h.Headers = []HeaderKeyPair{}
	for _, v := range headers[1:] {
		matches = re.FindStringSubmatch(v)
		h.Headers = append(h.Headers, HeaderKeyPair{Key: matches[1], Value: matches[2]})
	}

	return
}

func (h *ResponseHeader) ToString() string {
	buf := fmt.Sprintf("\"%s/%s %s\"", h.Protocol, h.Version, h.Code)
	for _, v := range h.Headers {
		buf = buf + " " + v.ToString()
	}
	return buf
}
