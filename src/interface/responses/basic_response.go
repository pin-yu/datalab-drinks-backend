package responses

// Response has two fields status and body. Response is an immutable object.
type Response struct {
	status int
	body   Body
}

// Body has two fields status message and  payload. Bodyis an immutable object.
type Body struct {
	StatusMessage string      `json:"status_message"`
	Payload       interface{} `json:"payload"`
}

// NewResponse is a constructor of response
func NewResponse(status int, statusMsg string, payload interface{}) *Response {
	body := Body{StatusMessage: statusMsg, Payload: payload}
	return &Response{status: status, body: body}
}

// Resolve returns statust and body
func (res *Response) Resolve() (int, Body) {
	return res.status, res.body
}
