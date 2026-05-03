package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestCovariance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		x        []float64
		y        []float64
		expected float64
	}{
		{[]float64{}, []float64{}, 0},
		{[]float64{1}, []float64{1}, 0},
		{[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}, 2.5},
		{[]float64{1, 2, 3, 4, 5}, []float64{5, 4, 3, 2, 1}, -2.5},
		{[]float64{10, 20, 30}, []float64{15, 25, 35}, 100.0},
		{[]float64{1, 2, 3}, []float64{1, 2}, 0},
	}
	for i, test := range tests {
		output := mathx.Covariance[float64, float64](test.x, test.y)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestCovariance_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(2.5), mathx.Covariance[float32, float32]([]float32{1, 2, 3, 4, 5}, []float32{1, 2, 3, 4, 5}))
}

func TestCorrelation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		x        []float64
		y        []float64
		expected float64
	}{
		{[]float64{}, []float64{}, 0},
		{[]float64{1}, []float64{1}, 0},
		{[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}, 0.9999999999999998},
		{[]float64{1, 2, 3, 4, 5}, []float64{5, 4, 3, 2, 1}, -0.9999999999999998},
		{[]float64{1, 2, 3, 4, 5}, []float64{2, 4, 6, 8, 10}, 0.9999999999999998},
		{[]float64{1, 2, 3}, []float64{1, 2}, 0},
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, 0},
	}
	for i, test := range tests {
		output := mathx.Correlation[float64, float64](test.x, test.y)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestCorrelation_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(1), mathx.Correlation[float32, float32]([]float32{1, 2, 3, 4, 5}, []float32{1, 2, 3, 4, 5}))
}
