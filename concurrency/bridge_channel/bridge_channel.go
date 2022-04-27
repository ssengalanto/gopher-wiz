package main

/*
	bridge channel destructure streams of channel of channels into a simple channel stream
*/
func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for {
			var stream <-chan interface{}
			select {
			case v, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = v
			case <-done:
				return
			}
			for val := range orDone(done, stream) {
				select {
				case c <- val:
				case <-done:
				}
			}
		}
	}()
	return c
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case stream <- v:
				case <-done:
				}
			}
		}
	}()
	return stream
}
