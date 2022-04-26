package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

/*
	The key thing to note here is how we’ve coupled the potential result with the
	potential error. This represents the complete set of possible outcomes from
	the goroutine checkStatus creates and allows our main goroutine to make
	decisions about what to do when errors occur. In broader terms, we’ve
	successfully separated the concerns of error handling from our producer
	goroutine. This is desirable because the goroutine which spawned the
	producer goroutine — in this case our main goroutine — has more context
	about the running program, and can make more intelligent decisions about
	what to do with errors.
*/

func main() {
	fn := func(done <-chan interface{}, urls ...string) <-chan Result {
		c := make(chan Result)
		go func() {
			defer close(c)
			for _, url := range urls {
				var res Result
				resp, err := http.Get(url)
				res = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case c <- res:
				}
			}
		}()
		return c
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://go.dev", "https://badhost"}
	for res := range fn(done, urls...) {
		if res.Error != nil {
			fmt.Printf("error: %v", res.Error)
			continue
		}
		fmt.Printf("Response: %v\n", res.Response.Status)
	}
}
