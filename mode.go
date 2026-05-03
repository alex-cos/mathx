package mathx

// Mode returns the mode (most frequent value) in the values.
// Returns the first occurring mode if there are ties.
func Mode[F Float](values []F) F {
	if len(values) == 0 {
		return 0
	}
	counts := make(map[F]int)
	maxCount := 0
	var mode F
	for _, v := range values {
		counts[v]++
		if counts[v] > maxCount {
			maxCount = counts[v]
			mode = v
		}
	}
	return mode
}
