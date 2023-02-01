package main

import "fmt"

type Present struct {
	V int64
	S int64
}

func (p *Present) Weight() int64 {
	return p.S
}

func (p *Present) Value() int64 {
	return p.S
}

func main() {
	src := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	fmt.Println(grabPresents(src, 6))
}

func grabPresents(src []Present, cap int) []Present {
	if cap <= 0 || len(src) <= 0 {
		return []Present{}
	}
	idx := Knapsack(src, int64(cap))
	result := make([]Present, 0, len(idx))
	for _, v := range idx {
		result = append(result, src[v])
	}
	return result
}

func Knapsack(items []Present, capacity int64) []int64 {

	values := make([][]int64, len(items)+1)
	for i := range values {
		values[i] = make([]int64, capacity+1)
	}

	keep := make([][]int, len(items)+1)
	for i := range keep {
		keep[i] = make([]int, capacity+1)
	}

	for i := int64(0); i < capacity+1; i++ {
		values[0][i] = 0
		keep[0][i] = 0
	}

	for i := 0; i < len(items)+1; i++ {
		values[i][0] = 0
		keep[i][0] = 0
	}

	for i := 1; i <= len(items); i++ {
		for c := int64(1); c <= capacity; c++ {

			itemFits := (items[i-1].Weight() <= c)
			if !itemFits {
				continue
			}

			maxValueAtThisCapacity := items[i-1].Value() + values[i-1][c-items[i-1].Weight()]
			previousValueAtThisCapacity := values[i-1][c]

			if itemFits && (maxValueAtThisCapacity > previousValueAtThisCapacity) {
				values[i][c] = maxValueAtThisCapacity
				keep[i][c] = 1
			} else {
				values[i][c] = previousValueAtThisCapacity
				keep[i][c] = 0
			}
		}
	}

	n := len(items)
	c := capacity
	var indices []int64

	for n > 0 {
		if keep[n][c] == 1 {
			indices = append(indices, int64(n-1))
			c -= items[n-1].Weight()
		}
		n--
	}

	return indices
}
