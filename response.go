package harlog

import "net/http"

// Response HAR representation of http response
type Response struct {

  /* TODO
     "cookies": [],
     "content": {},
     "redirectURL": "",
  */
  Status     int         `json:"status"`
  StatusText string      `json"statusText`
  Version    string      `json:"httpVersion"`
  Headers    http.Header `json:"headers"`
  BodySize   int64       `json:"bodySize"`
}

// NewResponse constructs a new HAR response
func NewResponse(r *http.Response) *Response {

  return &Response{
    Status:     r.StatusCode,
    StatusText: r.Status,
    Version:    r.Proto,
    Headers:    r.Header,
    BodySize:   r.ContentLength}
}
