package median

import (
	"cringe.com/mean"
	"sort"
)

func Median(sequence []int) float64 {
	sort.Ints(sequence)
	n := len(sequence)
	if n%2 == 0 {
		res := mean.Mean([]int{sequence[(n / 2)], sequence[(n/2)-1]})
		return res
	}
	res := sequence[(n / 2)]
	return float64(res)
}
