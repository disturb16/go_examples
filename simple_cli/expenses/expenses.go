package expenses

func Average(expns ...float32) float32 {
	return Sum(expns...) / float32(len(expns))
}

func Sum(expns ...float32) float32 {
	var sum float32

	for _, exp := range expns {
		sum += exp
	}

	return sum
}

func Max(expns ...float32) float32 {
	var max float32

	for _, exp := range expns {
		if exp > max {
			max = exp
		}
	}

	return max
}

func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}

	var min float32 = expns[0]

	for _, exp := range expns {
		if exp < min {
			min = exp
		}
	}

	return min
}
