package harlog

import (
  "encoding/json"
  "log"
)

const version = "1.2"

type creator struct {
  Name    string `json:"name"`
  Version string `json:"version"`
}

// HARLog contains requst / response lifecycle entries
type HARLog struct {
  Version string
  Creator creator
  Entries Entries
}

// Dump marshals the HARLog to JSON and logs it
func (h HARLog) Dump() {

  j, err := json.Marshal(h)
  if err == nil {
    log.Print(string(j))
  }
}

// NewHARLog constructs a new HAR logger
func NewHARLog() *HARLog {

  c := &creator{
    Name:    "harlog",
    Version: "1.0"}

  return &HARLog{
    Version: version,
    Entries: make([]*Entry, 0),
    Creator: *c}
}
