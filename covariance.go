package mathx

// Covariance returns the sample covariance between x and y.
func Covariance[N Number, F Float](x, y []N) F {
	n := len(x)
	if n != len(y) || n <= 1 {
		return 0.0
	}
	avgX := Average[N, F](x)
	avgY := Average[N, F](y)
	tmp := float64(0.0)
	for i := range n {
		tmp += (float64(x[i]) - float64(avgX)) * (float64(y[i]) - float64(avgY))
	}
	return F(tmp / float64(n-1))
}

// Correlation returns the Pearson correlation coefficient between x and y.
func Correlation[N Number, F Float](x, y []N) F {
	n := len(x)
	if n != len(y) || n <= 1 {
		return 0.0
	}
	stdX := StdDev[N, F](x)
	stdY := StdDev[N, F](y)
	if stdX == 0 || stdY == 0 {
		return 0.0
	}
	cov := Covariance[N, F](x, y)
	return cov / (stdX * stdY)
}
