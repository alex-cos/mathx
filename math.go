package mathx

import "math"

// Pow returns the value raised to the specified exponent.
func Pow[F Float](value, exponent F) F {
	return F(math.Pow(float64(value), float64(exponent)))
}

// PowAll returns each value raised to the specified exponent.
func PowAll[F Float](values []F, exponent F) []F {
	result := make([]F, len(values))
	for i, v := range values {
		result[i] = Pow(v, exponent)
	}
	return result
}

// Sqrt returns the square root of the value.
func Sqrt[F Float](value F) F {
	return F(math.Sqrt(float64(value)))
}

// SqrtAll returns the square root of each value.
func SqrtAll[F Float](values []F) []F {
	result := make([]F, len(values))
	for i, v := range values {
		result[i] = Sqrt(v)
	}
	return result
}

// Abs returns the absolute value.
func Abs[F Float](value F) F {
	if value < 0 {
		return -value
	}
	return value
}

// AbsAll returns the absolute value of each value.
func AbsAll[F Float](values []F) []F {
	result := make([]F, len(values))
	for i, v := range values {
		result[i] = Abs(v)
	}
	return result
}
