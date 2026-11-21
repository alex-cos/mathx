package mathx

import "math"

func niceStep[F Float](minVal, maxVal F, numIntervals int) F {
	rawStep := (maxVal - minVal) / F(numIntervals)
	rawStepf := float64(rawStep)
	if rawStepf == 0 {
		return rawStep
	}
	exp := math.Floor(math.Log10(rawStepf))
	pow10 := F(math.Pow(10, exp))
	normalized := rawStepf / math.Pow(10, exp)
	var niceNormalized F
	switch {
	case normalized <= 1:
		niceNormalized = 1
	case normalized <= 2:
		niceNormalized = 2
	case normalized <= 5:
		niceNormalized = 5
	default:
		niceNormalized = 10
	}
	return niceNormalized * pow10
}

// ComputeDistribution computes a distribution (histogram) of the values.
//
// The number of intervals is at most interval and never exceeds len(values).
// Interval bounds are rounded to the given precision (decimal places).
// Every interval is half-open [Inf, Sup), except the last one which also
// includes Sup, so the counts always add up to len(values).
//
// It returns an empty slice when values is empty or interval is not positive.
// When all values are equal, a single interval with Inf == Sup is returned.
func ComputeDistribution[F Float, I Integer](
	values []F, interval, precision int,
) []DistributionInterval[F, I] {
	d := min(interval, len(values))
	m := []DistributionInterval[F, I]{}
	if d <= 0 {
		return m
	}

	minF := Min(values)
	maxF := Max(values)
	p := math.Pow(10.0, float64(precision))
	round := func(v float64) F { return F(math.Round(v*p) / p) }

	if minF == maxF {
		return append(m, DistributionInterval[F, I]{
			Inf:   round(float64(minF)),
			Sup:   round(float64(maxF)),
			Count: I(len(values)),
		})
	}

	step := niceStep(minF, maxF, d)
	minimum := F(math.Floor(float64(minF)/float64(step))) * step
	maximum := F(math.Ceil(float64(maxF)/float64(step))) * step
	count := max(int(math.Round(float64(maximum-minimum)/float64(step))), 1)

	for k := range count {
		// Compute bounds from the index to avoid accumulating rounding errors.
		inf := minimum + F(k)*step
		sup := minimum + F(k+1)*step
		last := k == count-1
		if last {
			sup = maximum
		}
		nb := I(0)
		for _, val := range values {
			if val >= inf && (val < sup || (last && val <= sup)) {
				nb++
			}
		}
		m = append(m, DistributionInterval[F, I]{
			Inf:   round(float64(inf)),
			Sup:   round(float64(sup)),
			Count: nb,
		})
	}

	return m
}
