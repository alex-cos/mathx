package mathx_test

import (
	"testing"

	"github.com/alex-cos/mathx"
	"github.com/stretchr/testify/assert"
)

func TestMode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{1}, 1},
		{[]float64{1, 2, 3}, 1},
		{[]float64{1, 2, 2, 3}, 2},
		{[]float64{1, 2, 2, 3, 3, 3}, 3},
		{[]float64{5, 1, 3, 5, 2, 5}, 5},
		{[]float64{1.5, 2.5, 1.5, 2.5}, 1.5},
	}
	for i, test := range tests {
		output := mathx.Mode(test.input)
		assert.Equalf(t, test.expected, output, "test sample #%d", i)
	}
}

func TestMode_Float32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, float32(2), mathx.Mode([]float32{1, 2, 2, 3}))
}
