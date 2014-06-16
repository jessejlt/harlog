package harlog

import (
  "net/http"
  "strconv"
)

// Request HAR representation of http request
type Request struct {
  Method   string      `json:"method"`
  URL      string      `json:"url"`
  Version  string      `json:"httpVersion"`
  Query    string      `json:"queryString"`
  BodySize int         `json:"bodySize"`
  Headers  http.Header `json:"headers"`
  // TODO
  // "cookies": [],
  // "postData" : {},
}

// NewRequest constrcuts a new HAR request
func NewRequest(r *http.Request) *Request {

  bodyLen, err := strconv.Atoi(r.Header.Get("Content-Length"))
  if err != nil {
    bodyLen = 0
  }

  return &Request{
    Method:   r.Method,
    URL:      r.URL.String(),
    Version:  r.Proto,
    Query:    r.URL.RawQuery,
    BodySize: bodyLen,
    Headers:  r.Header}
}
