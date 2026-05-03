package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestMinMaxNormalize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{}, []float64{}},
		{[]float64{1}, []float64{1}},
		{[]float64{1, 2, 3}, []float64{0, 0.5, 1}},
		{[]float64{10, 20, 30}, []float64{0, 0.5, 1}},
		{[]float64{-10, 0, 10}, []float64{0, 0.5, 1}},
		{[]float64{5, 5, 5}, []float64{1, 1, 1}},
		{[]float64{100}, []float64{1}},
	}
	for i, test := range tests {
		output := mathx.MinMaxNormalize(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestMinMaxNormalize_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []float32{0, 0.5, 1}, mathx.MinMaxNormalize([]float32{1, 2, 3}))
}

func TestSumNormalize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{}, []float64{}},
		{[]float64{1, 2, 3}, []float64{1.0 / 6.0, 2.0 / 6.0, 3.0 / 6.0}},
		{[]float64{10, 20, 30}, []float64{1.0 / 6.0, 2.0 / 6.0, 3.0 / 6.0}},
		{[]float64{0, 0, 0}, []float64{0, 0, 0}},
	}
	for i, test := range tests {
		output := mathx.SumNormalize(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestSumNormalize_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []float32{1.0 / 6.0, 2.0 / 6.0, 3.0 / 6.0}, mathx.SumNormalize([]float32{1, 2, 3}))
}
