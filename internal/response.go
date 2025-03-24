package internal

import (
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode int
	Headers []Header
	Body string
}

func NewResponse(status int, body string) (*Response, error) {
	switch {
	case status < 100 || status > 599:
		return nil, errors.New("invalid status code")
	default:
		if body == ""{
			body = http.StatusText(status)
		}
		headers := []Header{}
		headers = append(headers, Header{"Content-Length", fmt.Sprintf("%d", len(body))})
		return &Response{
			StatusCode: status,
			Headers: headers,
			Body: body,
		}, nil
	}
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
