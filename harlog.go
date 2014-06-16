/*
Package harlog implements the HAR log version 1.2 spec http://www.softwareishard.com/blog/har-12-spec/

Create a new HAR log at the beginning of your request / response lifecycle then add new entries
each time you make additional requests within lifecycle.

TODO
Request: cookies, postData
Response: cookies, content, redirectURL

Example response
{
  "Version": "1.2",
  "Creator": {
    "name": "harlog",
    "version": "1.0"
  },
  "Entries": [
    {
      "startedDateTime": "Sun, 15 Jun 2014 18:12:41 PDT",
      "time": 0,
      "request": {
        "method": "GET",
        "url": "http://www.example.com/path/?param=value",
        "httpVersion": "HTTP/1.1",
        "queryString": "param=value",
        "bodySize": 0,
        "headers": {}
      },
      "response": {
        "status": 200,
        "StatusText": "200 OK",
        "httpVersion": "HTTP/1.0",
        "headers": {
          "Connection": [
            "close"
          ]
        },
        "bodySize": -1
      }
    }
  ]
}
*/
package harlog

import "encoding/json"

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
func (h HARLog) Dump() string {

  j, err := json.Marshal(h)
  if err == nil {
    return string(j)
  }

  return ""
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
