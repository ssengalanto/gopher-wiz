package main

/*
	Some handy generators
*/

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case c <- v:
				}
			}
		}
	}()
	return c
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case c <- valueStream:
			}
		}
	}()
	return c
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for {
			select {
			case <-done:
				return
			case c <- fn():
			}
		}
	}()
	return c
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for {
			for v := range valueStream {
				select {
				case <-done:
					return
				case c <- v.(string):
				}
			}
		}
	}()
	return c
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for {
			for v := range valueStream {
				select {
				case <-done:
					return
				case c <- v.(int):
				}
			}
		}
	}()
	return c
}
