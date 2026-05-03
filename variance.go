package mathx

import "math"

// Variance returns the sample variance of the values.
func Variance[N Number, F Float](values []N) F {
	nb := len(values)
	if nb <= 1 {
		return 0.0
	}
	avg := Average[N, F](values)
	tmp := float64(0.0)
	for _, val := range values {
		tmp += math.Pow(float64(val)-float64(avg), 2)
	}
	return F(tmp / float64(nb-1))
}

// StdDev returns the sample standard deviation of the values.
func StdDev[N Number, F Float](values []N) F {
	v := Variance[N, F](values)
	return F(math.Sqrt(float64(v)))
}
