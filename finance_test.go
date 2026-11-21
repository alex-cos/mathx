package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestDailyReturn(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    []float64
		expected []float64
	}{
		{
			[]float64{},
			[]float64{},
		},
		{
			[]float64{10, 20, 30, 40, 50},
			[]float64{1, 0.5, 0.33333333333333326, 0.25},
		},
		{
			[]float64{120, 119, 117, 114, 112, 113, 111, 110, 112, 115, 120, 117, 114, 105, 109, 113},
			[]float64{-0.008333333333333304, -0.01680672268907568, -0.02564102564102566,
				-0.01754385964912286, 0.008928571428571397, -0.017699115044247815,
				-0.009009009009009028, 0.018181818181818077, 0.02678571428571419,
				0.04347826086956519, -0.025000000000000022, -0.02564102564102566,
				-0.07894736842105265, 0.03809523809523818, 0.03669724770642202},
		},
	}

	for _, test := range tests {
		output := mathx.DailyReturn(test.input)
		assert.Equal(t, test.expected, output)
	}
}

func TestAnnualizedVolatility(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    []float64
		nbDays   float64
		expected float64
	}{
		{
			[]float64{},
			1.0,
			0.0,
		},
		{
			[]float64{0},
			1.0,
			0.0,
		},
		{
			[]float64{1},
			1.0,
			0.0,
		},
		{
			[]float64{1, 10},
			2,
			0,
		},
		{
			[]float64{10, 20, 3},
			7,
			3.4610330827658964,
		},
		{
			[]float64{100, 95, 105},
			10,
			0.3471789754539151,
		},
		{
			[]float64{100, 95, 105, 110, 150},
			30,
			0.9678296194285422,
		},
		{
			[]float64{120, 119, 117, 114, 112, 113, 111, 110, 112, 115, 120, 117, 114, 105, 109, 113},
			365.24,
			0.6226231434294442,
		},
	}

	for _, test := range tests {
		dr := mathx.DailyReturn(test.input)
		output := mathx.AnnualizedVolatility(dr, test.nbDays)
		assert.Equal(t, test.expected, output)
	}
}

func TestAnnualizedReturn(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    []float64
		nbDays   float64
		expected float64
	}{
		{
			[]float64{},
			1.0,
			0.0,
		},
		{
			[]float64{0},
			1.0,
			0.0,
		},
		{
			[]float64{1},
			1.0,
			0.0,
		},
		{
			[]float64{1, 10},
			2,
			99,
		},
		{
			[]float64{10, 20, 3},
			7,
			-0.9852114909473605,
		},
		{
			[]float64{100, 95, 105},
			10,
			0.27628156250000013,
		},
		{
			[]float64{100, 95, 105, 110, 150},
			30,
			19.92591432604223,
		},
		{
			[]float64{120, 119, 117, 114, 112, 113, 111, 110, 112, 115, 120, 117, 114, 105, 109, 113},
			365.24,
			-0.7685729286874,
		},
	}

	for _, test := range tests {
		dr := mathx.DailyReturn(test.input)
		output := mathx.AnnualizedReturn(dr, test.nbDays)
		assert.Equal(t, test.expected, output)
	}
}

func TestComputeDrawdownFromDR(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    []float64
		expected []float64
	}{
		{
			[]float64{},
			[]float64{},
		},
		{
			[]float64{-0.01506},
			[]float64{-0.015059999999999962},
		},
		{
			[]float64{
				-0.01506,
				0,
				0,
				0,
				0,
				-0.00778,
				0.006,
				-0.00083,
				0,
				-0.00534,
				0.0017,
				0,
			},
			[]float64{
				-0.015059999999999962,
				-0.015059999999999962,
				-0.015059999999999962,
				-0.015059999999999962,
				-0.015059999999999962,
				-0.022722833200000014,
				-0.016859170199200024,
				-0.017675177087934646,
				-0.017675177087934646,
				-0.02292079164228511,
				-0.021259756988076917,
				-0.021259756988076917,
			},
		},
	}

	for _, test := range tests {
		output := mathx.ComputeDrawdownFromDR(test.input)
		assert.Equal(t, test.expected, output)
	}
}

func TestComputeCumulPerfFromDR(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    []float64
		expected []float64
	}{
		{
			[]float64{},
			[]float64{},
		},
		{
			[]float64{
				-0.001427911915154,
				-0.006351502672168,
				0.010239945053225,
				0.016259262797984,
				0.004077068504435,
				0.03741830982241,
				-0.017181581066458,
				0.005005841671664,
				-0.02439444848084,
				0.007577508098963,
				0.026145199241012,
				-0.004280581471497,
				0.007757512937212,
				0.013344780736889,
				0.008925086162262,
			},
			[]float64{
				-0.001427911915154012,
				-0.007770345200977302,
				0.0023900319443450613,
				0.018688154899807907,
				0.022841416291990768,
				0.061114413305997095,
				0.04288278999299311,
				0.04810329612180131,
				0.022535394261959496,
				0.030283664493455875,
				0.057220636176397166,
				0.05269511710989616,
				0.060861413099816186,
				0.0750183760498595,
				0.08461300768211943,
			},
		},
	}

	for _, test := range tests {
		output := mathx.ComputeCumulPerfFromDR(test.input)
		assert.Equal(t, test.expected, output)
	}
}

func TestMaxDrawdown(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{-0.01}, -0.01},
		// Peak 1.0098 then trough 0.95931: -5%, not the -4.07% cumulated return.
		{[]float64{-0.01, 0.02, -0.05, 0.01}, -0.05},
		// Peak after three gains, then a -10% day.
		{[]float64{0.01, 0.02, 0.03, -0.10, 0.01}, -0.10},
		// Never recovers: drawdown from the starting level.
		{[]float64{-0.05, -0.03, -0.02}, -0.09693},
		// Only gains: no drawdown.
		{[]float64{0.01, 0.02, 0.03}, 0},
		// The deepest of two drawdowns wins.
		{[]float64{-0.02, 0.05, -0.10, 0.20, -0.05}, -0.10},
	}
	for i, test := range tests {
		output := mathx.MaxDrawdown(test.input)
		assert.InDeltaf(t, test.expected, output, 1e-12, "test sample #%d", i)
	}
}

func TestMaxDrawdown_MatchesDrawdownSeries(t *testing.T) {
	t.Parallel()
	returns := []float64{-0.01506, 0, 0.006, -0.00778, 0.0017, -0.00534, 0.02, -0.03}

	assert.Equal(t, mathx.Min(mathx.ComputeDrawdownFromDR(returns)), mathx.MaxDrawdown(returns))
	assert.LessOrEqual(t, mathx.MaxDrawdown(returns), 0.0)
}
