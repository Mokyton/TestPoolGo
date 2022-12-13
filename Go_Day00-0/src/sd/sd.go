package sd

import (
	"cringe.com/mean"
	"math"
)

func SD(sequence []int) float32 {
	var sum float64
	u := mean.Mean(sequence)
	l := len(sequence)
	for i := 0; i < l; i++ {
		sum += math.Pow(float64(sequence[i])-u, 2)
	}
	return float32(math.Sqrt(sum / float64(l)))
}
