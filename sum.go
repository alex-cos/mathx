package mathx

// Sum returns the sum of all values in the slice.
func Sum[N Number](values []N) N {
	sum := N(0.0)
	for _, val := range values {
		sum += val
	}
	return sum
}
