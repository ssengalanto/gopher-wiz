package main

import (
	"fmt"
	"time"
)

func main() {
	fmts := []string{
		time.Layout, time.ANSIC, time.UnixDate, time.RubyDate,
		time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123,
		time.RFC1123Z, time.RFC3339, time.RFC3339Nano, time.Kitchen,
		time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	}

	t := time.Now()
	for idx, val := range fmts {
		x := t.Format(val)
		fmt.Printf("%d: %s\n", idx, x)
	}
}
