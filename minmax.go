package mathx

// Maximum returns the maximum value among the given numbers.
func Maximum[N Number](nums ...N) N {
	if len(nums) == 0 {
		return 0
	}
	maximum := nums[0]
	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}

// Minimum returns the minimum value among the given numbers.
func Minimum[N Number](nums ...N) N {
	if len(nums) == 0 {
		return 0
	}
	minimum := nums[0]
	for _, num := range nums[1:] {
		if num < minimum {
			minimum = num
		}
	}

	return minimum
}

// Min returns the minimum value in the slice.
func Min[N Number](values []N) N {
	var minimum N
	if len(values) == 0 {
		return minimum
	}
	minimum = values[0]
	for _, val := range values {
		minimum = Minimum(minimum, val)
	}
	return minimum
}

// Max returns the maximum value in the slice.
func Max[N Number](values []N) N {
	var maximum N
	if len(values) == 0 {
		return maximum
	}
	maximum = values[0]
	for _, val := range values {
		maximum = Maximum(maximum, val)
	}
	return maximum
}

// Sign returns -1 if v is negative, otherwise returns +1 (including 0).
func Sign[N Number](v N) N {
	if v < 0 {
		return -1
	}
	return 1
}
