package internal

type Header struct{ Key, Value string }

type Request struct {
	Method  string
	Path    string
	Headers []struct{ Key, Value string }
	Body    string
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
