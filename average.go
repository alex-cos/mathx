package mathx

// Average returns the arithmetic mean of the values as type F.
func Average[N Number, F Float](values []N) F {
	nb := len(values)
	if nb == 0 {
		return 0.0
	}
	return F(float64(Sum(values)) / float64(nb))
}
