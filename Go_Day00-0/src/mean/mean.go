package mean

func Mean(sequence []int) float64 {
	var sum int
	for _, v := range sequence {
		sum += v
	}
	return float64(sum) / float64(len(sequence))
}
