package mathx

// MinMaxNormalize returns values normalized to [0, 1] using min-max scaling.
func MinMaxNormalize[F Float](values []F) []F {
	n := len(values)
	if n == 0 {
		return []F{}
	}
	if n == 1 {
		return []F{1}
	}
	minimum := Min(values)
	maximum := Max(values)
	if maximum == minimum {
		result := make([]F, n)
		for i := range result {
			result[i] = 1
		}
		return result
	}
	result := make([]F, n)
	for i, v := range values {
		result[i] = (v - minimum) / (maximum - minimum)
	}
	return result
}

// SumNormalize returns values normalized so their sum equals 1.
func SumNormalize[F Float](values []F) []F {
	n := len(values)
	if n == 0 {
		return []F{}
	}
	sum := Sum(values)
	if sum == 0 {
		result := make([]F, n)
		return result
	}
	result := make([]F, n)
	for i, v := range values {
		result[i] = v / sum
	}
	return result
}
