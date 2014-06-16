HARLog
---

Package harlog implements [HAR](http://www.softwareishard.com/blog/har-viewer/) log version [1.2](http://www.softwareishard.com/blog/har-12-spec/).

[godoc](https://godoc.org/github.com/jessejlt/harlog)

Use
---

```go
import (
  "github.com/jessejlt/harlog"
  "log"
  "net/http"  
)

client := &http.Client{}

function Handler(w http.ResponseWriter, r *http.Request) {

  // Create a new HAR log near beginning of your request handler
  har := NewHARLog()
  har.Entries.Add(r, w)

  // If you make any requests to other services, add them to your har log
  req, _ := http.NewRequest("GET", "http://example.com", nil)
  req.Header.Add("Authorization", "--")
  res, _ := client.Do(req)
  har.Entries.Add(req, res)

  // Dump the har log at the end of your request handler
  log.Printf("%+v", har.Dump())
}
```

Example Output
---

```json
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
```
