package harlog

import (
  "net/http"
  "testing"
)

func TestNewRequest(t *testing.T) {

  req, err := http.NewRequest("GET", "http://www.example.com/path/?param=value", nil)
  if err != nil {
    t.Error(err)
  }

  harReq := NewRequest(req)
  if harReq.Method != "GET" {
    t.Errorf("Expected method = GET, recieved %v", harReq.Method)
  }
  if harReq.Version != "HTTP/1.1" {
    t.Errorf("Expected version=1.1, recieved %v", harReq.Version)
  }
  if harReq.Query != "param=value" {
    t.Errorf("Expected query: param=value, recieved %v", harReq.Query)
  }
  if harReq.BodySize != 0 {
    t.Errorf("Expected BodySize=0, recieved %v", harReq.BodySize)
  }
}
