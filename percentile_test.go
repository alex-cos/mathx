package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestPercentile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		p        float64
		expected float64
	}{
		{[]float64{}, 0.5, 0.0},
		{[]float64{12.3}, 0.8, 12.3},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0.5, 5.5},
		{[]float64{3, 5, 7, 8, 9, 11, 13, 15}, 0.25, 6.5},
		{[]float64{35, 20, 50, 40, 15}, 0.4, 29.0},
		{[]float64{35, 20, 50, 40, 15}, 0.0, 15.0},
		{[]float64{35, 20, 50, 40, 15}, 1.0, 50.0},
		{[]float64{
			95.1772,
			95.1567,
			95.1937,
			95.1959,
			95.1442,
			95.0610,
			95.1591,
			95.1195,
			95.1065,
			95.0925,
			95.1990,
			95.1682,
		},
			0.9,
			95.19568,
		},
		{[]float64{
			-0.0434116445352399,
			-0.0898481334630058,
			-0.0280519076851348,
			-0.0003134796238244,
			-0.0163884208914077,
		},
			0.8,
			-0.013173432637891036,
		},
	}
	for i, test := range tests {
		output := mathx.Percentile(test.input, test.p)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestMedian(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{1}, 1},
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{5, 3, 1, 4, 2}, 3},
		{[]float64{10, 20, 30}, 20},
		{[]float64{-5, 0, 5}, 0},
	}
	for i, test := range tests {
		output := mathx.Median(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestMedian_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(2.5), mathx.Median([]float32{1, 2, 3, 4}))
}
