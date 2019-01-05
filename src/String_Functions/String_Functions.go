package main

import (
	"runtime"
	s "strings"
)
import "fmt"

var pn = fmt.Println
var pf = fmt.Printf

type point struct {
	x int
	y int
}

func main() {
	// string functions
	pn("Contains:  ", s.Contains("test", "es"))
	pn("Count:     ", s.Count("test", "t"))
	pn("HasPrefix: ", s.HasPrefix("test", "te"))
	pn("HasSuffix: ", s.HasSuffix("test", "st"))
	pn("Index:     ", s.Index("test", "e"))
	pn("Join:      ", s.Join([]string{"a", "b"}, "-"))
	pn("Repeat:    ", s.Repeat("a", 5))
	pn("Replace:   ", s.Replace("foo", "o", "0", -1))
	pn("Replace:   ", s.Replace("foo", "o", "0", 1))
	pn("Split:     ", s.Split("a-b-c-d-e", "-"))
	pn("ToLower:   ", s.ToLower("TEST"))
	pn("ToUpper:   ", s.ToUpper("test"))
	pn()

	pn("Len:       ", len("Hello"))
	pn("Char:      ", "golang"[3])

	// string format
	p := point{1, 2}
	pf("%v\n", p)
	pf("%+v\n", p)
	pf("%#v\n", p)
	pf("%T\n", p)

	pf("%t\n", true)
	pf("%d\n", 123)
	pf("%b\n", 123)
	pf("%c\n", 65)
	pf("%x\n", 123)
	pf("%f\n", 78.9)
	pf("%.2f\n", 78.9)

	pf("%e\n", 12340000.00)
	pf("%E\n", 12340000.00)

	pf(runtime.GOOS)
}
