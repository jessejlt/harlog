package harlog

import (
  "net/http"
  "time"
)

// Entry is a record of request / response lifecycle
type Entry struct {
  Started  string    `json:"startedDateTime"`
  Time     int       `json:"time"`
  Request  *Request  `json:"request"`
  Response *Response `json:"response"`
}

// Entries keeps track of request / response logs
type Entries []*Entry

// Add log entry
func (e *Entries) Add(req *http.Request) {

  entry := &Entry{
    Time:    0,
    Started: time.Now().Format(time.RFC1123),
    Request: NewRequest(req),
  }
  *e = append(*e, entry)
}
