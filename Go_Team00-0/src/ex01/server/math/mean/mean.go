package mean

func Mean(sequence []float64) float64 {
	var sum float64
	for _, v := range sequence {
		sum += v
	}
	return sum / float64(len(sequence))
}
