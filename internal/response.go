package internal

type Response struct {
	StatusCode int
	Headers []struct {Key, Value string}
	Body string
}
