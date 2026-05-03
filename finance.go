package mathx

import "math"

// DailyReturn computes the daily returns.
func DailyReturn[F Float](values []F) []F {
	nb := len(values)
	if nb == 0 {
		return []F{}
	}
	dailyReturns := make([]F, (nb - 1))
	for i := range nb - 1 {
		dailyReturns[i] = (values[i+1] / values[i]) - 1
	}
	return dailyReturns
}

// Volatility computes the volatility value for the given values.
func Volatility[F Float](values []F) F {
	return F(math.Sqrt(Variance[F, float64](values)))
}

// AnnualizedReturn computes the annualized return for the given values.
func AnnualizedReturn[F Float](values []F, nbDays F) F {
	nb := len(values)
	tmp := F(1.0)
	if nb == 0 {
		return 0.0
	}
	for _, val := range values {
		tmp *= (1 + val)
	}
	return F(math.Pow(float64(tmp), float64(nbDays)/float64(nb)) - 1.0)
}

// AnnualizedVolatility computes the annualized volatility for the given values.
func AnnualizedVolatility[F Float](values []F, nbDays F) F {
	return Volatility(values) * F(math.Sqrt(float64(nbDays)))
}

// ComputeDrawdownFromDR computes Drawdown from Daily Returns.
func ComputeDrawdownFromDR[N Number](dailyReturns []N) []N {
	nb := len(dailyReturns)
	maxDrawdown := N(0)
	if nb == 0 {
		return []N{}
	}

	drawdowns := make([]N, nb)
	cumulReturns := make([]N, nb)
	cumulReturns[0] = dailyReturns[0]
	for i := 1; i < nb; i++ {
		cumulReturns[i] = (1+cumulReturns[i-1])*(1+dailyReturns[i]) - 1
	}
	for i := range nb {
		maxDrawdown = Maximum(maxDrawdown, cumulReturns[i])
		drawdowns[i] = (1+cumulReturns[i])/(1+maxDrawdown) - 1
	}
	return drawdowns
}

// ComputeCumulPerfFromDR computes Cumulated Performance from Daily Returns.
func ComputeCumulPerfFromDR[N Number](dailyReturns []N) []N {
	nb := len(dailyReturns)
	if nb == 0 {
		return []N{}
	}
	perf := make([]N, nb)
	cumPerf := N(0.0)
	for i := range nb {
		cumPerf = (1+dailyReturns[i])*(1+cumPerf) - 1
		perf[i] = cumPerf
	}
	return perf
}

// MaxDrawdown returns the maximum drawdown from daily returns.
func MaxDrawdown[N Number](dailyReturns []N) N {
	nb := len(dailyReturns)
	if nb == 0 {
		return 0
	}
	cumulReturns := make([]N, nb)
	cumulReturns[0] = dailyReturns[0]
	for i := 1; i < nb; i++ {
		cumulReturns[i] = (1+cumulReturns[i-1])*(1+dailyReturns[i]) - 1
	}
	maxDrawdown := cumulReturns[0]
	for i := 1; i < nb; i++ {
		if cumulReturns[i] < maxDrawdown {
			maxDrawdown = cumulReturns[i]
		}
	}
	return maxDrawdown
}
