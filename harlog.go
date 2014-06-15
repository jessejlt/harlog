package harlog

import (
  "encoding/json"
  "log"
)

type Entry map[string][]string

func (e Entry) Add(key, value string) {
  e[key] = append(e[key], value)
}

type HARLog struct {
  Entry Entry
}

func (h HARLog) Dump() {

  j, err := json.Marshal(h.Entry)
  if err == nil {
    log.Print(string(j))
  }
}

func NewHARLog() *HARLog {

  return &HARLog{Entry: make(Entry)}
}
