package internal

import (
	"reflect"
	"testing"
)

func TestHTTPResponse(t *testing.T) {
	for name, tt := range map[string]struct {
		input string
		want  *Response
	}{
		"200 OK (no body)": {
			input: "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n",
			want: &Response{
				StatusCode: 200,
				Headers: []Header{
					{"Content-Length", "0"},
				},
			},
		},
		"404 Not Found (w/ body)": {
			input: "HTTP/1.1 404 Not Found\r\nContent-Length: 11\r\n\r\nHello World\r\n",
			want: &Response{
				StatusCode: 404,
				Headers: []Header{
					{"Content-Length", "11"},
				},
				Body: "Hello World",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ParseResponse(tt.input)
			if err != nil {
				t.Errorf("ParseResponse(%q) returned error: %v", tt.input, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseResponse(%q) = %#+v, want %#+v", tt.input, got, tt.want)
			}

			if got2, err := ParseResponse(got.String()); err != nil {
				t.Errorf("ParseResponse(%q) returned error: %v", got.String(), err)
			} else if !reflect.DeepEqual(got2, got) {
				t.Errorf("ParseResponse(%q) = %#+v, want %#+v", got.String(), got2, got)
			}

		})
	}
}
