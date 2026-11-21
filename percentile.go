package mathx

import (
	"math"
	"slices"
)

// Percentile returns the p-th percentile of the values using linear interpolation.
// p is expected in [0, 1] and is clamped to that range.
// It returns 0 for empty values and NaN if p is NaN.
func Percentile[F Float](values []F, p F) F {
	n := len(values)
	if n == 0 {
		return 0
	}
	if math.IsNaN(float64(p)) {
		return F(math.NaN())
	}
	if n == 1 {
		return values[0]
	}
	sorted := make([]F, n)
	copy(sorted, values)
	slices.Sort(sorted)
	if p <= 0 {
		return sorted[0]
	}
	if p >= 1 {
		return sorted[n-1]
	}
	rank := (float64(n - 1)) * float64(p)
	floor := math.Floor(rank)
	frac := F(rank - floor)
	pos := int(floor)
	if pos >= (n - 1) {
		return sorted[pos]
	}

	return sorted[pos] + frac*(sorted[pos+1]-sorted[pos])
}

// Median returns the median (50th percentile) of the values.
func Median[F Float](values []F) F {
	return Percentile(values, 0.5)
}
