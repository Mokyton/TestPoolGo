package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := generator(100)
	ch2 := generator(2)
	ch3 := generator(3333)
	out := multiplex(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}

func generator(data any) <-chan any {
	ch := make(chan any)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- data
		}
		close(ch)
	}()

	return ch
}

func multiplex(cs ...<-chan any) <-chan any {
	var wg sync.WaitGroup

	out := make(chan any)

	send := func(c <-chan any) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
