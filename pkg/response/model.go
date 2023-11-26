package response

type AppResponse struct {
	HttpStatus int    `json:"http_status"`
	Code       string `json:"code"`
	Messages   string `json:"messages"`
}

type ResponseMeta struct {
	TotalFiltered int `json:"total_filtered,omitempty"`
	PerPage       int `json:"per_page,omitempty"`
	CurrentPage   int `json:"current_page,omitempty"`
	Total         int `json:"total,omitempty"`
	TotalPage     int `json:"total_page,omitempty"`
}

type Response struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ResponsePagination struct {
	Code    string       `json:"code"`
	Data    interface{}  `json:"data"`
	Meta    ResponseMeta `json:"meta,omitempty"`
	Message string       `json:"message"`
	Errors  interface{}  `json:"errors,omitempty"`
}

type CaptureError struct {
	Trace   bool        `json:"type"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// "error": {
// "message": "Unknown path components: /6283821708285/messages",
// "type": "OAuthException",
// "code": 2500,
// "trace_id": "Amgr0jjZ0NZ-Ca2k5qo8HfV"
// }
type ErrorTrace struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Code    int         `json:"code"`
	TraceId string      `json:"trace_id,omitempty"`
	Trace   interface{} `json:"trace,omitempty"`
}
