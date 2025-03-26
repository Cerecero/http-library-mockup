package internal

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	StatusCode int
	Headers    []Header
	Body       string
}

func NewResponse(status int, body string) (*Response, error) {
	switch {
	case status < 100 || status > 599:
		return nil, errors.New("invalid status code")
	default:
		if body == "" {
			body = http.StatusText(status)
		}
		headers := []Header{}
		headers = append(headers, Header{"Content-Length", fmt.Sprintf("%d", len(body))})
		return &Response{
			StatusCode: status,
			Headers:    headers,
			Body:       body,
		}, nil
	}
}

func (resp *Response) WithHeader(key, value string) *Response {
	resp.Headers = append(resp.Headers, Header{AsTitle(key), value})
	return resp
}

func (resp *Response) WriteTo(w io.Writer) (n int64, err error) {
	printf := func(format string, args ...any) error {
		m, err := fmt.Fprintf(w, format, args...)
		n += int64(m)
		return err
	}
	if err := printf("HTTP/1.1 %d %s\r\n", resp.StatusCode, http.StatusText(resp.StatusCode)); err != nil {
		return n, err
	}

	for _, h := range resp.Headers {
		if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
			return n, err
		}
	}
	if err := printf("\r\n%s\r\n", resp.Body); err != nil {
		return n, err
	}
	return n, nil
}

var _ fmt.Stringer = (*Response)(nil) //Compile-time interface check
var _ encoding.TextMarshaler = (*Response)(nil)

func (resp *Response) String() string { b := new(strings.Builder); resp.WriteTo(b); return b.String() }
func (resp *Response) MarshalText() ([]byte, error) {
	b := new(bytes.Buffer)
	resp.WriteTo(b)
	return b.Bytes(), nil
}

func ParseResponse(raw string) (resp *Response, err error) {
    lines := splitLines(raw)
    log.Println(lines)

    first := strings.SplitN(lines[0], " ", 3)
    if !strings.Contains(first[0], "HTTP") {
        return nil, fmt.Errorf("malformed response: first line should contain HTTP version")
    }
    resp = new(Response)
    resp.StatusCode, err = strconv.Atoi(first[1])
    if err != nil {
        return nil, fmt.Errorf("malformed response: expected status code to be an integer, got %q", first[1])
    }
    if first[2] == "" || http.StatusText(resp.StatusCode) != first[2] {
        log.Printf("missing or incorrect status text for status code %d: expected %q, but got %q", resp.StatusCode, http.StatusText(resp.StatusCode), first[2])
    }
    var bodyStart int
    for i := 1; i < len(lines); i++ {
        log.Println(i, lines[i])
        if lines[i] == "" {
            bodyStart = i + 1
            break
        }
        key, val, ok := strings.Cut(lines[i], ": ")
        if !ok {
            return nil, fmt.Errorf("malformed response: header %q should be of form 'key: value'", lines[i])
        }
        key = AsTitle(key)
        resp.Headers = append(resp.Headers, Header{key, val})
    }
    resp.Body = strings.TrimSpace(strings.Join(lines[bodyStart:], "\r\n")) // recombine the body using normal newlines.
    return resp, nil
}
// HTTP response look like this
// <PROTOCOL/VERSION> <STATUS CODE> <STATUS MESSAGE>
// [<HEADER>: <VALUE>] (optional)
// [<HEADER>: <VALUE>]
// [<HEADER>: <VALUE>]

// [<RESPONSE BODY>] (optional)

// For example:
// HTTP/1.1 200 OK
// Content-Type: application/json
// Vary: Origin
// Vary: Access-Control-Request-Method
// Vary: Authorization
// Date: Mon, 24 Mar 2025 18:34:15 GMT
// Content-Length: 97
//
// {
//         "status": "available",
//         "system_info": {
//                 "environment": "development",
//                 "version": ""
//         }
// }
