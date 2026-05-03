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
func ComputeDistribution[F Float, I Integer](
	values []F, interval, precision int,
) []DistributionInterval[F, I] {
	minF := Min(values)
	maxF := Max(values)
	p := math.Pow(10.0, float64(precision))
	d := min(interval, len(values))
	m := []DistributionInterval[F, I]{}
	if d == 0 || len(values) == 0 {
		return m
	}

	step := niceStep(minF, maxF, d)
	minimum := F(math.Floor(float64(minF)/float64(step))) * step
	maximum := F(math.Ceil(float64(maxF)/float64(step))) * step

	for i := minimum; i < maximum; i += step {
		nb := I(0)
		inf := i
		sup := i + step
		for _, val := range values {
			valf64 := float64(val)
			if sup >= maximum {
				if valf64 >= float64(inf) && valf64 <= float64(sup) {
					nb++
				}
			} else {
				if valf64 >= float64(inf) && valf64 < float64(sup) {
					nb++
				}
			}
		}
		m = append(m, DistributionInterval[F, I]{
			Inf:   F(math.Round(float64(inf)*p) / p),
			Sup:   F(math.Round(float64(sup)*p) / p),
			Count: nb,
		})
	}

	return m
}
