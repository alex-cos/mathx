package mathx

// Average returns the arithmetic mean of the values as type F.
// The sum is accumulated in float64, so integer inputs cannot overflow.
func Average[N Number, F Float](values []N) F {
	nb := len(values)
	if nb == 0 {
		return 0.0
	}
	sum := 0.0
	for _, val := range values {
		sum += float64(val)
	}
	return F(sum / float64(nb))
}
