package internal

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Header struct{ Key, Value string }

type Request struct {
	Method  string
	Path    string
	Headers []Header
	Body    string
}

func NewRequest(method, path, host, body string) (*Request, error) {
	switch {
	case method == "":
		return nil, errors.New("missing required argument: method")
	case path == "":
		return nil, errors.New("missing required argument: path")
	case !strings.HasPrefix(path, "/"):
		return nil, errors.New("path muist start with /")
	case host == "":
		return nil, errors.New("missing required argument: host")
	default:
		headers := make([]Header, 2)
		headers[0] = Header{"Host", host}
		if body != "" {
			headers = append(headers, Header{"Content-Length", fmt.Sprintf("%d", len(body))})
		}
		return &Request{Method: method, Path: path, Headers: headers, Body: body}, nil
	}

}

func (r *Request) WithHeader(key, value string) *Request {
	r.Headers = append(r.Headers, Header{AsTitle(key), value})
	return r
}

func (r *Request) WriteTo(w io.Writer) (n int64, err error) {
	printf := func(format string, args ...any) error {
		m, err := fmt.Fprintf(w, format, args...)
		n += int64(m)
		return err
	}

	if err := printf("%s %s HTTP/1.1\r\n", r.Method, r.Path); err != nil {
		return n, err
	}

	for _, h := range r.Headers {
		if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
			return n, err
		}
	}

	printf("\r\n")
	err = printf("%s\r\n", r.Body)
	return n, err
}

// HTTP Request look like this
// <METHOD> <PATH> <PROTOCOL/VERSION>
// Host: <Host>
// [<HEADER>: <VALUE] (optional)
// [<HEADER>: <VALUE] (optional)
// [<HEADER>: <VALUE] (optional)

// [<REQUEST BODY>] (optional)

// For example
// GET /index.html HTTP/1.1
// Host: somewebsite.com

// NOTE: line breaks are windows style \r\n
