package mode

func Mode(sequence []int) int {
	set := make(map[int]int, 0)
	var mode int
	var max int
	for _, v := range sequence {
		set[v] += 1
	}

	for k, v := range set {
		if max < v {
			max = v
			mode = k
		}
	}

	for k, v := range set {
		if v == max && k < mode {
			mode = k
		}
	}
	return mode
}
