package harlog

import "testing"

// func TestNewLog(t *testing.T) {
//
//   har := NewHARLog()
//   har.Entries.Add()
//   har.Dump()
// }

// func TestDump(t *testing.T) {
//
//   har := NewHARLog()
//   har.Entries.Add()
//   har.Entries.Add()
//   har.Dump()
// }

// func TestNewLog(t *testing.T) {
//
//   har := NewHARLog()
//   entry := Entry{Time: 5}
//   har.Entries = append(har.Entries, entry)
//   har.Dump()
// }

func TestAddEntry(t *testing.T) {

  har := NewHARLog()
  har.Entries.Add()
  har.Entries.Add()
  har.Dump()
}
