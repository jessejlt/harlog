package harlog

import (
  "net/http"
  "testing"
)

func TestNewResponse(t *testing.T) {

  req, _ := http.NewRequest("GET", "http://www.example.com/path/?param=value", nil)
  res := &http.Response{
    Status:     "200 OK",
    StatusCode: 200,
    Proto:      "HTTP/1.0",
    ProtoMajor: 1,
    ProtoMinor: 0,
    Request:    req,
    Header: http.Header{
      "Connection": {"close"},
    },
    Close:         true,
    ContentLength: -1,
  }

  harRes := NewResponse(res)
  if harRes.Status != 200 {
    t.Errorf("Expected status=200, recieved=%v", harRes.Status)
  }
}
