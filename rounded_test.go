package mathx_test

import (
	"math"
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestRounded_FloatingPointNoise(t *testing.T) {
	t.Parallel()
	// 0.29*100 = 28.999999999999996 must not be truncated to 28.
	assert.Equal(t, 0.29, mathx.Rounded(0.29, 2, 0, false))
	assert.Equal(t, 0.57, mathx.Rounded(0.57, 2, 0, false))
	assert.Equal(t, 1.15, mathx.Rounded(1.15, 2, 0.05, false))
	// An increment such as 0.29 (scaled to 28.999999999999996) keeps its value.
	assert.Equal(t, 0.58, mathx.Rounded(0.6, 2, 0.29, false))
}

func TestRounded_UpperKeepsValuesOnTheGrid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input     float64
		precision int
		increment float64
		expected  float64
	}{
		{12.2, 2, 0.05, 12.2},
		{-12.2, 2, 0.05, -12.2},
		{12.25, 2, 0.05, 12.25},
		{0, 2, 0.05, 0},
		{5, 0, 1, 5},
		{0.29, 2, 0, 0.29},
		{0.29, 2, 0.01, 0.29},
		// Values just off the grid still move up.
		{12.201, 2, 0.05, 12.25},
		{1.2301, 2, 0, 1.24},
		{1.5, 0, 0, 2},
		{-1.5, 0, 0, -1},
	}
	for i, test := range tests {
		output := mathx.Rounded(test.input, test.precision, test.increment, true)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestRounded_LargeAndSpecialValues(t *testing.T) {
	t.Parallel()
	assert.InEpsilon(t, 1e300, mathx.Rounded(1e300, 2, 0.05, false), 1e-12)
	assert.Equal(t, 1e308, mathx.Rounded(1e308, 2, 0.05, false))
	assert.Equal(t, 1e17, mathx.Rounded(1e17, 2, 0.05, false))
	assert.True(t, math.IsNaN(mathx.Rounded(math.NaN(), 2, 0.05, false)))
	assert.True(t, math.IsInf(mathx.Rounded(math.Inf(1), 2, 0.05, false), 1))
	assert.True(t, math.IsInf(mathx.Rounded(math.Inf(-1), 2, 0.05, false), -1))
}

func TestRounded_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(12.2), mathx.Rounded(float32(12.18978), 2, float32(0.05), false))
}

func TestRounded(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input     float64
		precision int
		increment float64
		upper     bool
		expected  float64
	}{
		{12.18978, 2, 0.05, false, 12.2},
		{12.18978, 2, 0.05, true, 12.20},
		{112.3467986, 4, 0.0002, false, 112.3468},
		{112.3467986, 4, 0.0002, true, 112.3468},
		{34.79, 2, 0.05, true, 34.8},
		{201.3489, 2, 0.0001, true, 201.35},
		{201.3489, 2, 0.0001, false, 201.34},
		{-12.18978, 2, 0.05, false, -12.2},
		{-12.18978, 2, 0.05, true, -12.15},
		{-34.79, 2, 0.05, false, -34.8},
		{-0.29, 2, 0, false, -0.29},
		{-1.239, 2, 0, false, -1.24},
	}
	for i, test := range tests {
		output := mathx.Rounded(test.input, test.precision, test.increment, test.upper)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}
