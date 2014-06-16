package harlog

// Entry is a record of request / response lifecycle
type Entry struct {
  Time int
}

// Entries keeps track of request / response logs
type Entries []*Entry

// Add log entry
func (e *Entries) Add() {

  *e = append(*e, &Entry{Time: 5})
}
