package main

func tee(done, in <-chan interface{}) (_, _ <-chan interface{}) {
	x := make(chan interface{})
	y := make(chan interface{})

	go func() {
		defer close(x)
		defer close(y)

		for val := range orDone(done, in) {
			// shadowing
			var x, y = x, y
			// looping range of two to make sure to have equal write with x and y
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case x <- val:
					x = nil
				case y <- val:
					y = nil
				}
			}
		}
	}()
	return x, y
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
