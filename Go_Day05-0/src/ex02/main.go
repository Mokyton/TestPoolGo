package main

import (
	"container/heap"
	"errors"
	"ex02/myheap"
	"fmt"
	"log"
)

func getNCoolestPresents(src []myheap.Present, n int) ([]myheap.Present, error) {
	if n <= 0 {
		return nil, errors.New("n is equals to zero or negative ")
	} else if n > len(src) {
		return nil, errors.New("n is bigger than count of Presents")
	}
	result := make([]myheap.Present, 0, n)
	h := &myheap.PresentHeap{}

	heap.Init(h)
	for _, v := range src {
		heap.Push(h, v)
	}

	for i := 0; i < n; i++ {
		result = append(result, heap.Pop(h).(myheap.Present))
	}

	return result, nil
}

func main() {
	src := []myheap.Present{
		myheap.Present{5, 1},
		myheap.Present{4, 5},
		myheap.Present{3, 1},
		myheap.Present{5, 2},
	}
	rm, err := getNCoolestPresents(src, 2)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rm)
}
