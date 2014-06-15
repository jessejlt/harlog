package harlog

import "testing"

func TestNewLog(t *testing.T) {

  har := NewHARLog()
  har.Entry.Add("a", "b")
}

func TestDump(t *testing.T) {

  har := NewHARLog()
  har.Entry.Add("a", "b")
  har.Entry.Add("c", "d")
  har.Entry.Add("c", "e")
  har.Dump()
}
