package internal

type Response struct {
	StatusCode int
	Headers []struct {Key, Value string}
	Body string
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
