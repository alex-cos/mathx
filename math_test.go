package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    float64
		exponent float64
		expected float64
	}{
		{4, 2, 16},
		{2, 3, 8},
		{27, 1.0 / 3.0, 2.9999999999999996},
		{16, 0.5, 4},
		{0, 2, 0},
		{1, 0, 1},
		{2, 0, 1},
	}
	for i, test := range tests {
		output := mathx.Pow(test.input, test.exponent)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestPowAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		exponent float64
		expected []float64
	}{
		{[]float64{1, 2, 3}, 2, []float64{1, 4, 9}},
		{[]float64{4, 9, 16}, 0.5, []float64{2, 3, 4}},
		{[]float64{}, 2, []float64{}},
	}
	for i, test := range tests {
		output := mathx.PowAll(test.input, test.exponent)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    float64
		expected float64
	}{
		{4, 2},
		{9, 3},
		{16, 4},
		{0, 0},
		{2, 1.4142135623730951},
	}
	for i, test := range tests {
		output := mathx.Sqrt(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestSqrtAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{4, 9, 16}, []float64{2, 3, 4}},
		{[]float64{1, 4, 9}, []float64{1, 2, 3}},
		{[]float64{}, []float64{}},
	}
	for i, test := range tests {
		output := mathx.SqrtAll(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    float64
		expected float64
	}{
		{-5, 5},
		{5, 5},
		{0, 0},
		{-3.14, 3.14},
		{-0.001, 0.001},
	}
	for i, test := range tests {
		output := mathx.Abs(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestAbsAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{-1, 2, -3}, []float64{1, 2, 3}},
		{[]float64{-5, 0, 5}, []float64{5, 0, 5}},
		{[]float64{}, []float64{}},
	}
	for i, test := range tests {
		output := mathx.AbsAll(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}
