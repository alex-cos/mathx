package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestComputeDistribution(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input     []float64
		interval  int
		precision int
		expected  []mathx.DistributionInterval[float64, int32]
	}{
		{
			[]float64{},
			10,
			2,
			[]mathx.DistributionInterval[float64, int32]{},
		},
		{
			[]float64{0.1, 0.3, 0.5},
			10,
			3,
			[]mathx.DistributionInterval[float64, int32]{
				{0.0, 0.2, 1}, {0.2, 0.4, 1}, {0.4, 0.6, 1},
			},
		},
		{
			[]float64{-0.21, 0.33, 0.45, 0.12, -0.25, 0.78, 0.15, 0.08, 0.11, -0.76, 0.02},
			4,
			3,
			[]mathx.DistributionInterval[float64, int32]{
				{-1.0, -0.5, 1}, {-0.5, 0.0, 2}, {0.0, 0.5, 7}, {0.5, 1.0, 1},
			},
		},
		{
			[]float64{3.675, -2.897, 0.8674, -4.7107, -7.7936, -6.2536, -1.112, 3.3933, 11.2057,
				-2.0911, -4.6984, -9.0716, -9.6301, 4.0179, 4.5131, 4.0372, 3.1869, 4.6925,
				-2.1613, 1.1516, -1.8318, -0.2855, 2.2944, 5.3721, 9.4547},
			4,
			2,
			[]mathx.DistributionInterval[float64, int32]{
				{-10.0, 0.0, 12}, {0.0, 10.0, 12}, {10.0, 20.0, 1},
			},
		},
	}

	for i, test := range tests {
		output := mathx.ComputeDistribution[float64, int32](
			test.input,
			test.interval,
			test.precision,
		)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestComputeDistribution_Float32(t *testing.T) {
	t.Parallel()

	input := []float32{0.1, 0.3, 0.5}
	interval := 10
	precision := 3
	expected := []mathx.DistributionInterval[float32, int32]{
		{Inf: float32(0.0), Sup: float32(0.2), Count: int32(1)},
		{Inf: float32(0.2), Sup: float32(0.4), Count: int32(1)},
		{Inf: float32(0.4), Sup: float32(0.6), Count: int32(1)},
	}

	output := mathx.ComputeDistribution[float32, int32](input, interval, precision)
	assert.Equal(t, expected, output)
}

func TestComputeDistribution_EqualValues(t *testing.T) {
	t.Parallel()

	type di = mathx.DistributionInterval[float64, int32]
	assert.Equal(t, []di{{5, 5, 1}}, mathx.ComputeDistribution[float64, int32]([]float64{5}, 4, 2))
	assert.Equal(t, []di{{5, 5, 3}}, mathx.ComputeDistribution[float64, int32]([]float64{5, 5, 5}, 4, 2))
	assert.Equal(t, []di{{-2.5, -2.5, 2}}, mathx.ComputeDistribution[float64, int32]([]float64{-2.5, -2.5}, 4, 2))
	assert.Equal(t, []di{{0, 0, 2}}, mathx.ComputeDistribution[float64, int32]([]float64{0, 0}, 3, 2))
}

func TestComputeDistribution_InvalidInterval(t *testing.T) {
	t.Parallel()

	values := []float64{1, 2, 3}
	assert.Empty(t, mathx.ComputeDistribution[float64, int32](values, 0, 2))
	assert.Empty(t, mathx.ComputeDistribution[float64, int32](values, -1, 2))
	assert.NotNil(t, mathx.ComputeDistribution[float64, int32](values, -1, 2))
	assert.Empty(t, mathx.ComputeDistribution[float64, int32](nil, 4, 2))
}

func TestComputeDistribution_Invariants(t *testing.T) {
	t.Parallel()

	var values []float64
	for i := range 101 {
		values = append(values, float64(i)/100)
	}
	inputs := [][]float64{
		values,
		{3.675, -2.897, 0.8674, -4.7107, 11.2057, 9.4547},
		{1, 2, 3},
		{-3, 0, 3},
	}

	for n, input := range inputs {
		for _, interval := range []int{1, 2, 3, 4, 7, 10, 100} {
			bins := mathx.ComputeDistribution[float64, int32](input, interval, 6)
			if !assert.NotEmptyf(t, bins, "input #%d, interval %d", n, interval) {
				continue
			}

			var total int32
			for i, b := range bins {
				total += b.Count
				assert.Less(t, b.Inf, b.Sup)
				if i > 0 {
					// Contiguous intervals, no drift between bounds.
					assert.Equalf(t, bins[i-1].Sup, b.Inf, "input #%d, interval %d, bin %d", n, interval, i)
				}
			}
			assert.Equalf(t, int32(len(input)), total, "input #%d, interval %d", n, interval)
			assert.LessOrEqual(t, len(bins), min(interval, len(input))+1)
		}
	}
}
