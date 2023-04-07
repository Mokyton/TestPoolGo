package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := sleepSort([]int{10, 4, 0, 6, 9})
	for v := range ch {
		fmt.Println(v)
	}
}

func sleepSort(src []int) chan int {
	out := make(chan int, len(src))
	defer close(out)
	var wg sync.WaitGroup

	wg.Add(len(src))
	for _, v := range src {
		go func(v int) {
			defer wg.Done()
			time.Sleep(time.Duration(v) * time.Millisecond)
			out <- v
		}(v)
	}
	wg.Wait()
	return out
}
