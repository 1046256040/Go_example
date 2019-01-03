package main

import (
	"fmt"
	"time"
)

var pn = fmt.Println

func main() {
	now := time.Now()
	pn(now)

	nowUnix := now.Unix()
	pn(nowUnix)

	then := time.Date(2019, 1, 1, 15, 38, 10, 10, time.UTC)
	pn(then)

	pn(then.Year())
	pn(then.Month())
	pn(then.Day())
	pn(then.Hour())
	pn(then.Minute())
	pn(then.Second())
	pn(then.Nanosecond())
	pn(then.Location())

	pn(then.Before(now))
	pn(then.After(now))
	pn(then.Equal(now))

	diff := now.Sub(then)
	pn(diff)
	pn(diff.Hours())
	pn(diff.Minutes())
	pn(diff.Seconds())
	pn(diff.Nanoseconds())

	pn(then.Add(diff))
	pn(then.Add(-diff))
}
