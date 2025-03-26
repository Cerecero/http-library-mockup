package internal

import (
	"reflect"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	for name, tt := range map[string]struct {
		input string
		want  Request
	}{
		"GET (no body)": {
			input: "GET / HTTP/1.1\r\nHost: www.example.com\r\n\r\n",
			want: Request{
				Method: "GET",
				Path:   "/",
				Headers: []Header{
					{"Host", "www.example.com"},
				},
			},
		},
		"POST (w/ body)": {
			input: "POST / HTTP/1.1\r\nHost: www.example.com\r\nContent-Length: 11\r\n\r\nHello World\r\n",
			want: Request{
				Method: "POST",
				Path:   "/",
				Headers: []Header{
					{"Host", "www.example.com"},
					{"Content-Length", "11"},
				},
				Body: "Hello World",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ParseRequest(tt.input)
			if err != nil {
				t.Errorf("ParseRequest(%q) returned error: %v", tt.input, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequest(%q) = %#+v, want %#+v", tt.input, got, tt.want)
			}
			// test that the request can be written to a string and parsed back into the same request.
			got2, err := ParseRequest(got.String())
			if err != nil {
				t.Errorf("ParseRequest(%q) returned error: %v", got.String(), err)
			}
			if !reflect.DeepEqual(got, got2) {
				t.Errorf("ParseRequest(%q) = %+v, want %+v", got.String(), got2, got)
			}

		})
	}
}
