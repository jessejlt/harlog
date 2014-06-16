package harlog

import (
  "net/http"
  "testing"
)

func TestAddEntry(t *testing.T) {

  req, _ := http.NewRequest("GET", "http://www.example.com/path/?param=value", nil)
  har := NewHARLog()
  har.Entries.Add(req)
  har.Dump()
}
