package mathx

import "math"

// Rounded returns the value rounded to the given precision and increment.
func Rounded[F Float](
	value F, precision int, increment F, upper bool,
) F {
	p := math.Pow(10.0, float64(precision))
	res := int64(float64(value) * p)
	i := int64(float64(increment) * p)
	if i > 0 {
		res = (res / i) * i
	}
	if upper || float64(value)*p > float64(res)+float64(i)/2.0 {
		res += i
	}

	return F(float64(res) / p)
}
