package main

import (
	"fmt"
	"time"
)

/*
	Or-channel combines one or more done channels into a single done channel which closes if any of its component channels close
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	fn := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		fn(2*time.Hour),
		fn(5*time.Minute),
		fn(1*time.Second),
		fn(1*time.Hour),
		fn(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
