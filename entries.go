package harlog

import "time"

// Entry is a record of request / response lifecycle
type Entry struct {
  Started string  `json:"startedDateTime"`
  Time    int     `json:"time"`
  Request Request `json:"request"`
}

// Entries keeps track of request / response logs
type Entries []*Entry

// Add log entry
func (e *Entries) Add() {

  entry := &Entry{
    Time:    0,
    Started: time.Now().Format(time.RFC1123)}
  *e = append(*e, entry)
}
