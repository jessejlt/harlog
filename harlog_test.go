package harlog

import (
  "net/http"
  "testing"
)

func makerr() (*http.Request, *http.Response) {

  req, _ := http.NewRequest("GET", "http://www.example.com/path/?param=value", nil)
  resp := &http.Response{
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

  return req, resp
}

func TestNewLog(t *testing.T) {

  har := NewHARLog()
  har.Dump()
}

func TestAddEntry(t *testing.T) {

  req, resp := makerr()
  har := NewHARLog()
  har.Entries.Add(req, resp)
}

func TestDump(t *testing.T) {

  req, resp := makerr()
  har := NewHARLog()
  har.Entries.Add(req, resp)
  har.Dump()
}
