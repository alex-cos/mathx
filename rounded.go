package mathx

import "math"

// snapUlps is how many units in the last place a scaled value may be away
// from an integer to still be considered that integer.
const snapUlps = 8

// snapToInteger returns the nearest integer when x is within a few ulps of it.
// It absorbs binary floating-point noise such as 0.29*100 = 28.999999999999996.
func snapToInteger(x float64) float64 {
	r := math.Round(x)
	ax := math.Abs(x)
	if math.Abs(x-r) <= snapUlps*(math.Nextafter(ax, math.Inf(1))-ax) {
		return r
	}
	return x
}

// Rounded returns the value rounded to the given precision and increment.
//
// The value is first scaled by 10^precision. When the scaled increment is
// positive, the result is the nearest multiple of that increment (ties go
// down). When the scaled increment is zero, the value is rounded down (towards
// negative infinity) to the given precision.
//
// When upper is true, the value is instead rounded up (ceiling) to the next
// multiple of the increment, or to the given precision when the scaled
// increment is zero. A value already on that grid is returned unchanged.
//
// NaN, infinities and values too large to scale are returned unchanged.
func Rounded[F Float](
	value F, precision int, increment F, upper bool,
) F {
	v := float64(value)
	p := math.Pow(10.0, float64(precision))
	scaled := snapToInteger(v * p)
	if math.IsNaN(scaled) || math.IsInf(scaled, 0) || p == 0 {
		return value
	}

	res := math.Floor(scaled)
	i := math.Floor(snapToInteger(float64(increment) * p))
	if i > 0 {
		res = math.Floor(res/i) * i
	}
	switch {
	case upper:
		// Ceiling: only move up when the value is not already on the grid.
		if res < scaled {
			res += max(i, 1)
		}
	case scaled > res+i/2.0:
		res += i
	}

	return F(res / p)
}
